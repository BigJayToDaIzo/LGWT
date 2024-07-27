package sink

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment 3 times returns 3", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()
		assertCounter(t, c, 3)
	})
	t.Run("increments concurrency safe", func(t *testing.T) {
		want := 1000
		c := NewCounter()
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, c, want)
	})
}

func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()
	if c.Value() != want {
		t.Errorf("got %d, want %d", c.Value(), want)
	}
}
