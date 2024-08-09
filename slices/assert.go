package main

import "testing"

func AssertEqual[U comparable](t *testing.T, got, want U) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[U comparable](t *testing.T, got, want U) {
	t.Helper()
	if got == want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got false, want true")
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got true, want false")
	}
}
