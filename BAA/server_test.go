package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// stub up a player store for testing
// this should help us uncover other store types
// file written to disk, or a database somewhere for example
type StubPlayerStore struct {
	// seems a good choice for in memory store
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	scores, ok := s.scores[name]
	if !ok {
		return 0, false
	}
	return scores, true
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Oboe")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseCode(t, response.Code, http.StatusNotFound)
	})
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

// Helpers
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func assertResponseCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response code is wrong, got %d, want %d", got, want)
	}
}
