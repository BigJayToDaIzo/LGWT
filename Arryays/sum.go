package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	lenOfNums := len(numsToSum)
	sums := make([]int, lenOfNums)
	for i, n := range numsToSum {
		sums[i] = Sum(n)
	}
	return sums
}

// clips head and sums the rest
func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int
	for _, n := range tailsToSum {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			tail := n[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
