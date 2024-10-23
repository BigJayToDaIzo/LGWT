package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2+2=4", func(t *testing.T) {
		sum := Adder(2, 2)
		want := 4
		if sum != want {
			t.Errorf("sum %d want %d", sum, want)
		}
	})
	t.Run("3+5=8", func(t *testing.T) {
		sum := Adder(3, 5)
		want := 8
		if sum != want {
			t.Errorf("sum %d want %d", sum, want)
		}
	})
}

// Ensure we're writing Examples for the documentation!
// ExampleFuncName() and a // Ouptut: example output at the end!

func ExampleAdder() {
	sum := Adder(2, 2)
	fmt.Println(sum)
	// Output: 4
}
