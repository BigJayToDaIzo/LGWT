package main

import (
	"bytes"
	"testing"
)

// Test suite nees 0 second duration sleep timer, lets give it one!
type SpySleeper struct {
	// Make sure our countdown is called the appropriate # of times
	calls int
	// write the method for our struct
}

func (s *SpySleeper) Sleep() {
	// what would you have your little spy do instead of sleep?
	// what's the phone bill gonna be like?
	s.calls++
}

func TestCountdown(t *testing.T) {
	// t.Run("prints 3", func(t *testing.T) {
	// 	// bytes.Buffer in the test suite
	// 	b := &bytes.Buffer{}
	// 	Countdown(b)
	//
	// 	got := b.String()
	// 	want := "3"
	//
	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// })
	t.Run("prints countdown", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &SpySleeper{}
		Countdown(b, s)
		got := b.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		// ensure timer called 3x
		if s.calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", s.calls)
		}
	})
}
