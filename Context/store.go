package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// I like to think of context as 'a list of shid that could go wrong'
// then it must choose what to do if any of that list goes wrong
// First step is to check for completeness
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			fmt.Println("Error: ", err)
			return // todo: log error

		}
		fmt.Fprint(w, data)
	}
}

// saved for documentation (without rolling back commits)
// this is where the test suite took us for simplicity
// before we refactore the context into the store
// 	// grab context
// 	ctx := r.Context()
// 	// create channel for race
// 	data := make(chan string, 1)
// 	// start data race
// 	go func() {
// 		data <- store.Fetch()
// 	}()
// 	// determine the winner
// 	select {
// 	case d := <-data:
// 		fmt.Fprint(w, d)
// 	case <-ctx.Done():
// 		store.Cancel()
// 	}
