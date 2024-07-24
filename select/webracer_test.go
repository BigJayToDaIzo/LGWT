package webracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of fastest one",
		func(t *testing.T) {
			slowURL, fastURL := returnServerSet(20, 0)
			defer slowURL.Close()
			defer fastURL.Close()
			want := fastURL.URL
			got, e := Racer(slowURL.URL, fastURL.URL)
			if e != nil {
				t.Errorf("got an unexpected error %v", e)
			}
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	t.Run("returns an error if both servers times out (11ms)",
		func(t *testing.T) {
			slowURL, fastURL := returnServerSet(12, 12)
			defer slowURL.Close()
			defer fastURL.Close()
			_, err := ConfigurableRacer(fastURL.URL, slowURL.URL, 11*time.Millisecond)
			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		})
	t.Run("does not error if at least one server returns within timeout",
		func(t *testing.T) {
			slowURL, fastURL := returnServerSet(12, 0)
			defer slowURL.Close()
			defer fastURL.Close()
			_, err := ConfigurableRacer(fastURL.URL, slowURL.URL, 11*time.Millisecond)
			if err != nil {
				t.Error("unexpected error", err)
			}
			want := fastURL.URL
			got, _ := Racer(fastURL.URL, slowURL.URL)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
}

func returnServerSet(s, f int) (slow, fast *httptest.Server) {
	return makeDelayedServer(s), makeDelayedServer(f)

}

func makeDelayedServer(delay int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
