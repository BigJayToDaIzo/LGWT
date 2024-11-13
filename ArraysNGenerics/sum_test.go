package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkEquality := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum slices and return a slice of their sums", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}
		checkEquality(t, got, want)
	})
	t.Run("sum 3 slices of 3", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
		want := []int{6, 15, 24}
		checkEquality(t, got, want)
	})
	t.Run("sum slices of varying len", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7}, []int{8, 9})
		want := []int{6, 22, 17}
		checkEquality(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	// tighten up scope on helper funcs to keep api privacy strong
	checkEquality := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum tails of slices and return a slice of their sums", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}
		checkEquality(t, got, want)
	})
	t.Run("returns []int{0} for size 0 array or nil", func(t *testing.T) {
		got := SumAllTails(nil)
		want := []int{0}
		checkEquality(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concat strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concat, ""), "abc")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
