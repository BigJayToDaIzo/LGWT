package main

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
