package main

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(trans []Transaction, name string) float64 {
	var balance float64
	for _, t := range trans {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
