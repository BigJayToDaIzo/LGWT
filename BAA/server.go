package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	// these are both only declared in test stubs for now
	// TDD will show us the way
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// First we read from the store (Get)
	// We inject player/score/status into the case methods where needed
	// player := strings.TrimPrefix(r.URL.Path, "/players/")
	// score, ok := p.store.GetPlayerScore(player)
	// Then we write to the store (Post)
	// if r.Method == http.MethodPost {
	// 	// lets give everything StatusAccepted and see how that goes.
	// 	// this will foce us to consider updating other tests to add these assertions
	// 	w.WriteHeader(http.StatusAccepted)
	// 	return
	// }
	// this becomes a prime refactor target, switch on r.Method
	// this helps us abstract the process win logic out of the ServeHTTP method
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r)
	case http.MethodPost:
		p.processWin(w)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	// We need the player
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, score)

}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	// we begin by sending status accepted for everything
	// remember how this helps us find other weaknesses in our tests!
	w.WriteHeader(http.StatusAccepted)
	// Next up, sad path of win processing
	// of course we nest w.WriteHeader(http.StatusAccepted) within the happy branch
	// playstore is required here
	p.store.RecordWin("Pepper")
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
