package main

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

func main() {
	// we need a router now!
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	// log.Fatal(http.ListenAndServe(":5000", server))
}
