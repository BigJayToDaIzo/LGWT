package slices

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(tr []Transaction, name string) float64 {
	var balance float64
	for _, trans := range tr {
		if trans.From == name {
			balance -= trans.Sum
		}
		if trans.To == name {
			balance += trans.Sum
		}
	}
	return balance
}

func SumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func Reduce[A any](collection []A, f func(A, A) A, initValue A) A {
	var result = initValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func SumSlices(sums []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(sums, add, 0)
}

func SumSliceTails(sums ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, SumSlices(tail))
		}
	}
	return Reduce(sums, sumTail, []int{})
}
