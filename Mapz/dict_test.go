package main

import (
	"testing"
)

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		dict := Dict{"test": "testing testicle"}
		dict.Delete("test")
		err, _ := dict.Search("test")
		assertStrings(t, err.Error(), ErrDictBroken.Error())
	})
	t.Run("delete non-existent key", func(t *testing.T) {
		dict := Dict{}
		err := dict.Delete("test")
		assertStrings(t, err.Error(), ErrDictInadequate.Error())
	})
}

func TestSearch(t *testing.T) {
	t.Run("existent map key", func(t *testing.T) {
		want := "testing testicle"
		dict := Dict{"test": want}
		_, result := dict.Search("test")
		assertStrings(t, result, want)
	})
	t.Run("non existent map key", func(t *testing.T) {
		dict := Dict{"test": "testing testicle"}
		err, _ := dict.Search("testkl")
		assertStrings(t, err.Error(), ErrDictBroken.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("add non existent key", func(t *testing.T) {
		dict := Dict{}
		want := "testing testicle"
		_ = dict.Add("test", want)
		_, result := dict.Search("test")
		assertStrings(t, result, want)
	})
	t.Run("add existing key", func(t *testing.T) {
		dict := Dict{"p": "splitch splatch"}
		err := dict.Add("p", "splatch splitch")
		assertStrings(t, err.Error(), ErrDictTooFull.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		dict := Dict{"test": "testing testicle"}
		want := "testing testis"
		_ = dict.Update("test", want)
		assertStrings(t, dict["test"], want)
	})
	t.Run("update non-existent key", func(t *testing.T) {
		dict := Dict{}
		err := dict.Update("test", "testing testis")
		assertStrings(t, err.Error(), ErrDictInadequate.Error())
	})
}
