package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Walking..."},
			[]string{"Walking..."},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Bertha", "St. Louis"},
			[]string{"Bertha", "St. Louis"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Bertha", 29},
			[]string{"Bertha"},
		},
		{
			"struct with nested fields",
			Person{
				"Bertha",
				Profile{29, "St. Louis"},
			},
			[]string{"Bertha", "St. Louis"},
		},
		{
			"pointer to a struct",
			&Person{
				"Bertha",
				Profile{29, "St. Louis"},
			},
			[]string{"Bertha", "St. Louis"},
		},
		{
			"slices",
			[]Profile{
				{29, "St. Louis"},
				{34, "Chicago"},
			},
			[]string{"St. Louis", "Chicago"},
		},
		{
			"arrays",
			[2]Profile{
				{29, "St. Louis"},
				{34, "Chicago"},
			},
			[]string{"St. Louis", "Chicago"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
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
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{29, "St. Louis"}
			aChannel <- Profile{34, "Chicago"}
			close(aChannel)
		}()
		var got []string
		want := []string{"St. Louis", "Chicago"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{29, "St. Louis"}, Profile{34, "Chicago"}
		}
		var got []string
		want := []string{"St. Louis", "Chicago"}
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
