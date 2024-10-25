package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}
	Greet(&b, "Boi!")
	got := b.String()
	want := "Oy! Boi!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
