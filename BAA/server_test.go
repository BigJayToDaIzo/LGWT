package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// stub up a player store for testing
// this should help us uncover other store types
// file written to disk, or a database somewhere for example
type StubPlayerStore struct {
	// seems a good choice for in memory store
	scores   map[string]int
	league   []Player
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

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		[]string{},
	}
	server := NewPlayerServer(&store)
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
		nil,
	}
	server := NewPlayerServer(&store)

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
			nil,
		}
		player := "Pepper"
		server := NewPlayerServer(&store)
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

// NEW USER STORY
// customer wants a /league endpoint
func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)
	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		// NOTE: NOW WE BRING Jason ON THE SCENE!
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		assertResponseCode(t, response.Code, http.StatusOK)
	})
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		server := NewPlayerServer(&StubPlayerStore{nil, wantedLeague, nil})
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response.Body)

		assertResponseCode(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
	})

}

// Helpers
func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
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

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}
