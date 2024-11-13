package main

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numsToSum ...[]int) []int {
	sumAll := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(x))
	}
	return Reduce(numsToSum, sumAll, []int{})
}

// clips head and sums the rest
func SumAllTails(tailsToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		tail := x[1:]
		return append(acc, Sum(tail))
	}
	return Reduce(tailsToSum, sumTail, []int{})
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result

}
