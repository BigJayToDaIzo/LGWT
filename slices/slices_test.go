package slices

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiplyFn := func(x, y int) int {
			return x * y
		}
		assertCorrect(t, Reduce([]int{1, 2, 3}, multiplyFn, 1), 6)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concatenateFn := func(x, y string) string {
			return x + y
		}
		assertCorrect(t, Reduce([]string{"a", "b", "c"}, concatenateFn, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	// (t, expected, actual)
	assertCorrect(t, newBalanceFor(riya), 200)
	assertCorrect(t, newBalanceFor(chris), 0)
	assertCorrect(t, newBalanceFor(adil), 175)
}

func assertCorrect[A comparable](t *testing.T, actual, expected A) {
	if actual != expected {
		t.Helper()
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
