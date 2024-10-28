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
}
