package slices

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		assertCorrectness(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		assertCorrectness(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	trans := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	assertCorrectness(t, BalanceFor(trans, "Riya"), 100)
	assertCorrectness(t, BalanceFor(trans, "Chris"), -75)
	assertCorrectness(t, BalanceFor(trans, "Adil"), -25)
}

//	func TestSumSlice(t *testing.T) {
//		sumSliceTests := []struct {
//			Slice []int
//			want  int
//			desc  string
//		}{
//			{[]int{1, 2, 3, 4, 5}, 15, "sum of 1 to 5"},
//			{[]int{2, 3, 4, 5}, 14, "sum of 2 to 5"},
//			{[]int{}, 0, "empty slice"},
//		}
//		for _, tt := range sumSliceTests {
//			fmt.Println(tt.desc)
//			assertCorrectness(t, tt.want, SumSlice(tt.Slice))
//		}
//	}
//
//	func TestSumSlices(t *testing.T) {
//		sumSlicesTests := []struct {
//			Slices [][]int
//			want   int
//			desc   string
//		}{
//			{[][]int{{1, 2, 3}, {4, 5, 6}}, 21, "sum of 1 to 6 (2 slices)"},
//			{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 45, "sum of 1 to 9 (3 slices)"},
//			{[][]int{{}}, 0, "empty slice"},
//		}
//		for _, tt := range sumSlicesTests {
//			fmt.Println(tt.desc)
//			assertCorrectness(t, tt.want, SumSlices(tt.Slices...))
//		}
//	}
//
//	func TestSumSliceTails(t *testing.T) {
//		testSumSliceTails := []struct {
//			Slices [][]int
//			want   int
//			desc   string
//		}{
//			{[][]int{{1, 2, 3}, {4, 5, 6}}, 9, "sum of slice tails 3 and 6"},
//			{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 18, "sum of slice tails 3, 6, and 9"},
//			{[][]int{{}}, 0, "empty slice"},
//		}
//		for _, tt := range testSumSliceTails {
//			fmt.Println(tt.desc)
//			assertCorrectness(t, tt.want, SumSliceTails(tt.Slices...))
//		}
//	}
func assertCorrectness[A comparable](t *testing.T, expected, actual A) {
	if actual != expected {
		t.Helper()
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
