package main

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
	t.Run("find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Yoshi"},
			{Name: "Danielle"},
			{Name: "Chris"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris"})
	})
}

func AssertTrue(t *testing.T, b bool) {
	if !b {
		t.Errorf("expected true but got false")
	}
}
