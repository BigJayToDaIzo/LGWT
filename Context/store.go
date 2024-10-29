package main

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// grab context
		ctx := r.Context()
		// create channel for race
		data := make(chan string, 1)
		// start data race
		go func() {
			data <- store.Fetch()
		}()
		// determine the winner
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
