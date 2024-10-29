package main

import (
	"sync"
	"testing"
)

// Paraphrasing from LGWT/Sync/synk.go:
// * Use channels when PASSING OWNERSHIP of data
// * Use mutexex for managing state
func TestCounter(t *testing.T) {
	assertCounter := func(t testing.TB, got *Counter, want int) {
		t.Helper()
		if got.Value() != want {
			t.Errorf("got %d, want %d", got.Value(), want)
		}
	}
	t.Run("inc 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, &counter, 3)
	})
	t.Run("safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}
