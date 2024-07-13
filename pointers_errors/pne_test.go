package pne

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("stringer", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		got := wallet.Balance().String()
		want := "10 BTC"
		assertStringCorrectness(t, got, want)
	})
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertCorrectness(t, got, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertCorrectness(t, got, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(100))
		var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
		errorCheck(t, err, ErrInsufficientFunds)
		got := wallet.Balance()
		assertCorrectness(t, got, Bitcoin(20))
	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func errorCheck(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, got nil")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertCorrectness(t *testing.T, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertStringCorrectness(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
