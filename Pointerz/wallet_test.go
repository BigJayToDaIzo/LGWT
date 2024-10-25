package main

import (
	"testing"
)

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		// test is fatally executed here
		t.Fatal("Error assertion failed. None returned.")
		// that code down there doesn't get run, cuz this test is dead, remember?
	}
	// grab error string for comparison
	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got.String(), want.String())
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)
		if err != nil {
			t.Fatal("Appropriate withdraw failed.")
		}
		assertBalance(t, wallet, 10)
	})
	// Overdraw
	t.Run("Overdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)
		assertBalance(t, wallet, 20)
		assertError(t, err, ErrInsufficientCoinz.Error())
	})
}
