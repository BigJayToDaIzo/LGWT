package main

import (
	"net/http/httptest"
	"testing"
)

// We write integration test for InMemoryStore even tho it's only temporary
// this integration test can be reused for local disk stores, database stores, etc
// often in memory stores are a valuable option for APIs anyway.  Might not delete later.
func TestRecordingWinsAndRetreivingThem(t *testing.T) {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertResponseBody(t, response.Body.String(), "3")
}
