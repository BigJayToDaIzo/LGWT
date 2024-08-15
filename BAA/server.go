package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	jsonContentType = "application/json"
)

type PlayerStore interface {
	// NOTE: below has been updated with in_memory_player_store.go refactor
	// these are both only declared in test stubs for now
	// TDD will show us the way SPOILER: TDD showed us the way to above refactor
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	// Embedding the router
	// THIS embedding is how we retired the ServeHTTP archive notes/code
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.Handler = router

	return p
}

// router handler abstraction goes here
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	// We need the player
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	// we begin by sending status accepted for everything
	// remember how this helps us find other weaknesses in our tests!
	// Next up, sad path of win processing
	// of course we nest w.WriteHeader(http.StatusAccepted) within the happy branch
	// playstore is required here
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
