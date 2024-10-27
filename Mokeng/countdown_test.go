package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// Test suite nees 0 second duration sleep timer, lets give it one!
type SpyCountdownOps struct {
	// Make sure our countdown is called the appropriate # of times
	Calls []string
	// write the method for our struct
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOps) Sleep() {
	// what would you have your little spy do instead of sleep?
	// what's the phone bill gonna be like?
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints countdown", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &SpyCountdownOps{}
		Countdown(b, s)
		got := b.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("order of ops", func(t *testing.T) {
		s := &SpyCountdownOps{}
		// pass spy instead of buffer for write ops capture
		Countdown(s, s)
		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}
		if !reflect.DeepEqual(s.Calls, want) {
			t.Errorf("got %v want %v", s.Calls, want)
		}

	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("configurable sleeper", func(t *testing.T) {
		duration := 5 * time.Second
		durationSpy := &SpyTime{}
		s := &ConfigurableSleeper{duration, durationSpy.Sleep}
		s.Sleep()
		if durationSpy.durationSlept != duration {
			t.Errorf("Should have slept for %v but slept for %v", duration, durationSpy.durationSlept)
		}
	})
}
