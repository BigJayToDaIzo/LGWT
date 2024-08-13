package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	return "10"
}

// below func made obsolete by graduating to a struct with our own ServeHTTP method
// keeping around for notes and documentation
// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	// Parse the request url for player name
// 	// NOTE: TDD has brought us to ROUTING before we have
// 	// even thought about the data structure giving us much
// 	// smaller incremental steps to the solution
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
// 	fmt.Fprint(w, GetPlayerScore(player))
// }
