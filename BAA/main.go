package main

import (
	"log"
	"net/http"
)

// First we tackle in memory
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	return 123, true
}

// Then we tackle reading/writing to disk

// THEN we tackle SQLite3

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
