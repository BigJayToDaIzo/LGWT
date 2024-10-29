package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertCancelation(want bool) {
	s.t.Helper()
	if s.cancelled != want {
		s.t.Errorf("got %v, want %v", s.cancelled, want)
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store uncancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertCancelation(false)
		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}

	})
	t.Run("tells store to cancel work if request is cancelled",
		func(t *testing.T) {
			data := "hello, world"
			store := &SpyStore{response: data}
			svr := Server(store)

			// build basic request to /
			request := httptest.NewRequest(http.MethodGet, "/", nil)

			// build context that will fire cancelFn after 5ms
			cancellingCtx, cancelFn := context.WithCancel(request.Context())
			// set timer to wait then call cancel function
			time.AfterFunc(5*time.Millisecond, cancelFn)
			// pass context to request
			request = request.WithContext(cancellingCtx)

			// pass request to server and wait the 5ms for cancelFn to send
			// out signal to all goroutines that depend on it
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)
			store.assertCancelation(true)
		})

}
