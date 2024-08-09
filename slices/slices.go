package main

func SumSlices(sums []int) int {
	addFn := func(acc, x int) int { return acc + x }
	return Reduce(sums, addFn, 0)
}

func SumSliceTails(sums ...[]int) []int {
	// Sum all but the first element
	sumTailFn := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, SumSlices(tail))
		}
	}
	// This confused me at first. The function above is being declared
	// then passed as an argument to Reduce.
	return Reduce(sums, sumTailFn, []int{})
}
