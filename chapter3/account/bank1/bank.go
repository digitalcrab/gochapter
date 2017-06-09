package bank

// this package version is not thread safe

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}
