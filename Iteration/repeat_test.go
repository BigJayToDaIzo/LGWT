package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
