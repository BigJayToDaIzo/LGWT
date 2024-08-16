package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// We write integration test for InMemoryStore even tho it's only temporary
// this integration test can be reused for local disk stores, database stores, etc
// often in memory stores are a valuable option for APIs anyway.  Might not delete later.
func TestRecordingWinsAndRetreivingThem(t *testing.T) {
	// we now retire our temporary In Memory Store.  Let us slot in the FileStore now
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	// things broke here because of an empty file.  Now we write the test to ensure
	// we properly handle an empty file.
	store := assertFileStore(t, database)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertResponseCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		assertResponseCode(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := []Player{{"Pepper", 3}}
		assertLeague(t, got, want)
	})

}
