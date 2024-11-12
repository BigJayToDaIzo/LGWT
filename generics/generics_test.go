package generics

import "testing"

func TestGenericStack(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s_int := new(Stack[int])
		s_str := new(Stack[string])
		if !s_int.IsEmpty() || !s_str.IsEmpty() {
			t.Error("expected empty stacks")
		}
	})
	t.Run("push stack", func(t *testing.T) {
		s_int := new(Stack[int])
		s_str := new(Stack[string])
		s_int.Push(1)
		s_str.Push("a")
		if s_int.IsEmpty() || s_str.IsEmpty() {
			t.Error("expected non-empty stacks")
		}
	})
	t.Run("peek stack", func(t *testing.T) {
		s_int := new(Stack[int])
		s_str := new(Stack[string])
		s_int.Push(1)
		s_str.Push("a")
		i_val, _ := s_int.Peek()
		s_val, _ := s_str.Peek()
		AssertEqual(t, i_val, 1)
		AssertEqual(t, s_val, "a")
	})
	t.Run("peek empty stack throws error", func(t *testing.T) {
		s_int := new(Stack[int])
		s_str := new(Stack[string])
		_, err_int := s_int.Peek()
		_, err_str := s_str.Peek()
		if err_int == nil || err_str == nil {
			t.Error("expected error thrown")
		}
		AssertEqual(t, err_int.Error(), "cannot peek empty stack")
		AssertEqual(t, err_str.Error(), "cannot peek empty stack")
	})
	t.Run("pop stack", func(t *testing.T) {
		s_int := new(Stack[int])
		s_str := new(Stack[string])
		s_int.Push(1)
		s_str.Push("a")
		got_int, _ := s_int.Pop()
		got_str, _ := s_str.Pop()
		AssertEqual(t, got_int, 1)
		AssertEqual(t, got_str, "a")
		AssertEqual(t, s_int.IsEmpty(), true)
		AssertEqual(t, s_int.IsEmpty(), true)
	})
	t.Run("pop empty stack throws error", func(t *testing.T) {
		s_int := new(Stack[int])
		s_str := new(Stack[string])
		_, err_int := s_int.Pop()
		_, err_str := s_str.Pop()
		if err_int == nil || err_str == nil {
			t.Error("expected error thrown")
		}
		AssertEqual(t, err_int.Error(), "cannot pop empty stack")
		AssertEqual(t, err_str.Error(), "cannot pop empty stack")
	})
}

func TestStackOfInts(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := new(StackOfInts)
		got := s.IsEmpty()
		want := true
		AssertEqual(t, got, want)

	})
	t.Run("push stack", func(t *testing.T) {
		s := new(StackOfInts)
		AssertEqual(t, s.IsEmpty(), true)
		s.Push(1)
		AssertEqual(t, s.IsEmpty(), false)

	})
	t.Run("peek stack", func(t *testing.T) {
		s := new(StackOfInts)
		s.Push(1)
		got, _ := s.Peek()
		AssertEqual(t, got, 1)
	})
	t.Run("peek empty stack", func(t *testing.T) {
		s := new(StackOfInts)
		_, err := s.Peek()
		if err == nil {
			t.Errorf("expected error thrown")
		}
		AssertEqual(t, err.Error(), "cannot peek empty stack")
	})
	t.Run("pop stack", func(t *testing.T) {
		s := new(StackOfInts)
		s.Push(1)
		s.Push(2)
		got, _ := s.Pop()
		AssertEqual(t, got, 2)
		AssertEqual(t, s.IsEmpty(), false)
	})
	t.Run("pop empty stack throws error", func(t *testing.T) {
		s := new(StackOfInts)
		_, err := s.Pop()
		if err == nil {
			t.Errorf("expected error but got none")
		}
		AssertEqual(t, err.Error(), "cannot pop empty stack")
	})
}

func TestStackOfStrings(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := new(StackOfStrings)
		if !s.IsEmpty() {
			t.Errorf("expected empty stack")
		}
	})
	t.Run("push stack", func(t *testing.T) {
		s := new(StackOfStrings)
		s.Push("a")
		AssertEqual(t, s.IsEmpty(), false)
	})
	t.Run("peek stack", func(t *testing.T) {
		s := new(StackOfStrings)
		s.Push("a")
		got, _ := s.Peek()
		if got != "a" {
			t.Errorf("got %s, want a", got)
		}
	})
	t.Run("pop empty stack throws error", func(t *testing.T) {
		s := new(StackOfStrings)
		_, err := s.Pop()
		if err == nil {
			t.Errorf("expected error thrown")
		}
		AssertEqual(t, err.Error(), "cannot pop empty stack")

	})
	t.Run("peek empty stack throws error", func(t *testing.T) {
		s := new(StackOfStrings)
		_, err := s.Peek()
		if err == nil {
			t.Errorf("expected error thrown")
		}
		AssertEqual(t, err.Error(), "cannot peek empty stack")
	})
}

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "a", "a")
		AssertNotEqual(t, "a", "b")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("did not want matching values %+v & %+v", got, want)
	}
}
