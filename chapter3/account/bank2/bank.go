package bank

// buffered channel as a semaphore

var (
	balance int
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
