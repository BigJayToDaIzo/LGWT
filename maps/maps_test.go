package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		m := Dictionary{"foo": "bar"}
		got, _ := m.Search("foo")
		want := "bar"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		m := Dictionary{"foo": "bar"}
		_, err := m.Search("baz")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("already existing key", func(t *testing.T) {
		m := Dictionary{"foo": "bar"}
		err := m.Add("foo", "baz")
		assertError(t, err, ErrKeyExists)
		assertStrings(t, m["foo"], "bar") // Add attempt did not mutate map
	})
	t.Run("add new key", func(t *testing.T) {
		m := Dictionary{"foo": "bar"}
		err := m.Add("baz", "qux")
		if err != nil {
			t.Error("unexpected error", err)
		}
		assertStrings(t, m["baz"], "qux")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		m := Dictionary{"foo": "bar"}
		err := m.Update("foo", "baz")
		if err != nil {
			t.Error("unexpected error", err)
		}
		assertStrings(t, m["foo"], "baz")
	})
	t.Run("update non-existing key returns error", func(t *testing.T) {
		m := Dictionary{}
		err := m.Update("foo", "bar")
		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		m := Dictionary{"foo": "bar"}
		err := m.Delete("foo")
		if err != nil {
			t.Error("unexpected error", err)
		}
		assertStrings(t, m["foo"], "")
	})
	t.Run("delete non-existing key returns error", func(t *testing.T) {
		m := Dictionary{}
		err := m.Delete("foo")
		assertError(t, err, ErrDelKeyNotFound)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("got nil expected error")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
