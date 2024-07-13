package slices

func SumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func SumSlices(sums ...[]int) int {
	sum := 0
	for _, s := range sums {
		if len(s) == 0 {
			sum += 0
		} else {
			sum += SumSlice(s)
		}
	}
	return sum
}

func SumSliceTails(sums ...[]int) int {
	tailSum := 0
	for _, s := range sums {
		if len(s) == 0 {
			tailSum += 0
		} else {
			tailSum += s[len(s)-1]
		}
	}
	return tailSum
}
