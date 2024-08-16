package main

import (
	"log"
	"net/http"
	"os"
)

// below replaced with more permanent in_memory_player_store.go
// First we tackle in memory
// func (i *NewInMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
// 	return 123, true
// }
// func (i *InMemoryPlayerStore) RecordWin(name string) {}

// Then we tackle reading/writing to disk
// THEN we tackle Postgres
// TODO: Expand this boilerplate REST-API to interface with Postgres
// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#write-enough-code-to-make-it-pass-6

const dbFileName = "game.db.json"

func main() {
	// we need a router now!
	// NOTE: In memory now being retired for file io store
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	// log.Fatal(http.ListenAndServe(":5000", server))
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store := &FileSystemPlayerStore{db}
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
