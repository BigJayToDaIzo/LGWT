package geriatrics

import (
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])
		// check for empty stack
		AssertTrue(t, myStackOfInts.IsEmpty())

		// push an element
		myStackOfInts.Push(1)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// peek top element
		peek, ok := myStackOfInts.Peek()
		AssertEqual(t, peek, 1)
		AssertTrue(t, ok)

		// pop element
		popped, ok := myStackOfInts.Pop()
		AssertEqual(t, popped, 1)
		AssertTrue(t, ok)
		AssertTrue(t, myStackOfInts.IsEmpty())

		// ensure pop fails gracefully when stack is empty
		// int -> (int, bool)
		popped, ok = myStackOfInts.Pop()
		AssertEqual(t, popped, 0)
		AssertFalse(t, ok)

		// ensure peek fails gracefully when stack is empty
		// int -> (int, bool)
		popped, ok = myStackOfInts.Peek()
		AssertEqual(t, popped, 0)
		AssertFalse(t, ok)

	})
	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("hello")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		peeked, ok := myStackOfStrings.Peek()
		AssertEqual(t, peeked, "hello")
		AssertTrue(t, ok)

		popped, ok := myStackOfStrings.Pop()

		AssertEqual(t, popped, "hello")
		AssertTrue(t, ok)
		AssertTrue(t, myStackOfStrings.IsEmpty())

		popped, ok = myStackOfStrings.Pop()
		AssertEqual(t, popped, "")
		AssertFalse(t, ok)

		peeked, ok = myStackOfStrings.Peek()
		AssertEqual(t, peeked, "")
		AssertFalse(t, ok)

	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("Didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got false want true")
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got true want false")
	}
}
