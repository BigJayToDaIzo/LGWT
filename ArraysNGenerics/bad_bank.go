package main

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From, To string
	Sum      float64
}

// give our reducer a declarative name!
// woo how readable!
func applyTransaction(account Account, tr Transaction) Account {
	if tr.From == account.Name {
		account.Balance -= tr.Sum
	}
	if tr.To == account.Name {
		account.Balance += tr.Sum
	}
	return account
}

func NewBalanceFor(account Account, trans []Transaction) Account {
	return Reduce(trans, applyTransaction, account)
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: amount}
}

// refactor added account struct
// func BalanceFor(trans []Transaction, name string) float64 {
// 	adjustBalance := func(currentBalance float64, tr Transaction) float64 {
// 		if tr.From == name {
// 			currentBalance -= tr.Sum
// 		}
// 		if tr.To == name {
// 			currentBalance += tr.Sum
// 		}
// 		return currentBalance
// 	}
// 	return Reduce(trans, adjustBalance, 0.0)
// }
