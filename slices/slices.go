package slices

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, balance float64) Transaction {
	return Transaction{from.Name, to.Name, balance}
}

func NewBalanceFor(acc Account, trans []Transaction) Account {
	applyTrans := func(a Account, t Transaction) Account {
		if t.From == a.Name {
			a.Balance -= t.Sum
		}
		if t.To == a.Name {
			a.Balance += t.Sum
		}
		return a
	}
	return Reduce(trans, applyTrans, acc)
}

func BalanceFor(tr []Transaction, name string) float64 {
	adjBalanceFn := func(currentBal float64, tr Transaction) float64 {
		if tr.From == name {
			return currentBal - tr.Sum
		}
		if tr.To == name {
			return currentBal + tr.Sum
		}
		return currentBal
	}
	return Reduce(tr, adjBalanceFn, 0.0)
}

func Reduce[A, B any](collection []A, f func(B, A) B, initValue B) B {
	var result = initValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func SumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

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
