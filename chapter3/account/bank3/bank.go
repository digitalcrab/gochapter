package bank

import "sync"

// Use mutes to guard the balance

var (
	balance int
	mu      = sync.RWMutex{}
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	b := balance
	mu.RUnlock()
	return b
}
