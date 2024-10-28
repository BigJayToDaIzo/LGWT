package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("test Walk", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"nested fields",
				Person{
					"@DN",
					Profile{33, "London"},
				},
				[]string{"@DN", "London"},
			},
			{
				"pointer stuffs",
				&Person{
					"@DN",
					Profile{33, "London"},
				},
				[]string{"@DN", "London"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{34, "The Stl"},
				},
				[]string{"London", "The Stl"},
			},
		}
		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})
				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bu",
			"Baz": "Bozz",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bu")
		assertContains(t, got, "Bozz")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "@DN"}
			aChannel <- Profile{34, "The Stl"}
			close(aChannel)
		}()
		var got []string
		want := []string{"@DN", "The Stl"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "@DN"}, Profile{34, "The Stl"}
		}
		var got []string
		want := []string{"@DN", "The Stl"}
		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
