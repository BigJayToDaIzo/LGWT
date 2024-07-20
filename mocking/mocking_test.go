package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyCountdownOps{}
		Countdown(buffer, spy)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOps{}
		Countdown(spySleepPrinter, spySleepPrinter)
		spy_want := []string{
			write, sleep, write, sleep, write, sleep, write}
		if !reflect.DeepEqual(spy_want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", spy_want, spySleepPrinter.Calls)
		}
	})
	t.Run("configurable sleep timer", func(t *testing.T) {

	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("test sleep duration", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()
		if spyTime.durationSlept != sleepTime {
			t.Errorf("slept for %v want %v", spyTime.durationSlept, sleepTime)
		}
	})
}
