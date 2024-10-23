package main

import (
	"fmt"
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

// ExampleRepeat is an example of how to use the Repeat function
func ExampleRepeat() {
	got := Repeat("a", 5)
	fmt.Println(got)
	// Output: aaaaa
}

// BenchmarkRepeat is a benchmark test for the Repeat function
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
