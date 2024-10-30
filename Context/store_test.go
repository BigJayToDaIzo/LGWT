package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	// wheee! context is into the store!
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got canceled: Ctx10msTIMEOUT")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}

	// preserving for posterity to see how context design evolved
	// retired when we abstract context to store
	// time.Sleep(100 * time.Millisecond)
	// return s.response, nil
}

func (s *SpyStore) Cancel() {
	// s.cancelled = true
}

func (s *SpyStore) assertCancelation() {
	s.t.Helper()
	// if s.cancelled != want {
	// 	s.t.Errorf("")
	// }
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store uncancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		// store.assertCancelation(false)
		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}

	})
	t.Run("tells store to cancel work if request is cancelled",
		func(t *testing.T) {
			data := "hello, world"
			store := &SpyStore{response: data, t: t}
			svr := Server(store)

			// build basic request to /
			request := httptest.NewRequest(http.MethodGet, "/", nil)

			// build context that will fire cancelFn after 5ms
			// here shid that could go wrong is taking longer than 5ms
			// in production this would obvi be closer to 3m but we don't want
			// to slow the test suite by 3m so we do 5ms in the mock instead
			cancellingCtx, cancelFn := context.WithCancel(request.Context())
			// set timer to wait then call cancel function
			time.AfterFunc(5*time.Millisecond, cancelFn)
			// pass context to request
			request = request.WithContext(cancellingCtx)

			// pass request to server and wait the 5ms for cancelFn to send
			// out signal to all goroutines that depend on it
			response := &SpyResponseWriter{}

			svr.ServeHTTP(response, request)
			if response.written {
				t.Error("store should have canceled request, response should not be written")
			}
		})

}
