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
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	scores, ok := s.scores[name]
	if !ok {
		return 0, false
	}
	return scores, true
}

// RECORD THE WIN!
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
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

// Now we store the score in the store! For shore!
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		// should return 404 on red run, add Pepper to the store
		// now 200 instead of 202.  How do we get to 202?
		// back to ServeHTTP it would seem!
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusAccepted)
	})
	t.Run("it records wins when POST", func(t *testing.T) {
		// reset fresh server and store
		store := StubPlayerStore{
			map[string]int{},
			nil,
		}
		player := "Pepper"
		server := &PlayerServer{&store}
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not record correct winner, got %q, want %q", store.winCalls[0], player)
		}
	})
}

// Helpers
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
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
