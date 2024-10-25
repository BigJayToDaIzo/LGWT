package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientCoinz = errors.New("woopsies! git moar coinz lolz!")

func (w *Wallet) Withdraw(coin Bitcoin) error {
	if coin > w.balance {
		return ErrInsufficientCoinz
	}
	w.balance -= coin
	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
