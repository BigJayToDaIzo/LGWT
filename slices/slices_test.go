package slices

import "testing"

func TestSumSlice(t *testing.T) {
	t.Run("TestSumSlice", func(t *testing.T) {
		sum := []int{1, 2, 3, 4, 5}
		expected := 15
		actual := SumSlice(sum)
		assertCorrectness(t, expected, actual)
	})
	t.Run("TestSumSlice2", func(t *testing.T) {
		sum := []int{2, 3, 4, 5}
		expected := 14
		actual := SumSlice(sum)
		assertCorrectness(t, expected, actual)
	})
	t.Run("empty slice works", func(t *testing.T) {
		sum := []int{}
		expected := 0
		actual := SumSlice(sum)
		assertCorrectness(t, expected, actual)
	})
}

func TestSumSlices(t *testing.T) {
	t.Run("TestSumSlices", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5, 6}
		expected := 21
		actual := SumSlices(s1, s2)
		assertCorrectness(t, expected, actual)
	})
	t.Run("variadic works", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5, 6}
		s3 := []int{7, 8, 9}
		expected := 45
		actual := SumSlices(s1, s2, s3)
		assertCorrectness(t, expected, actual)
	})
	t.Run("empty slice works", func(t *testing.T) {
		s1 := []int{}
		expected := 0
		actual := SumSlices(s1)
		assertCorrectness(t, expected, actual)
	})
}

func TestSumSliceTails(t *testing.T) {
	t.Run("TestSumSliceTails", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5, 6}
		expected := 9
		actual := SumSliceTails(s1, s2)
		assertCorrectness(t, expected, actual)
	})
	t.Run("variable size slice works", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{5, 6}
		s3 := []int{9}
		expected := 18
		actual := SumSliceTails(s1, s2, s3)
		assertCorrectness(t, expected, actual)
	})
	t.Run("empty slice works", func(t *testing.T) {
		s1 := []int{}
		expected := 0
		actual := SumSliceTails(s1)
		assertCorrectness(t, expected, actual)
	})
}

func assertCorrectness(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Helper()
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
