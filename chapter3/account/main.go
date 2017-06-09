package main

import (
	"fmt"
	// try bank1, bank2 and bank3
	bank "gochapter/chapter3/account/bank1"
	"time"
)

func main() {
	t1 := time.Tick(2 * time.Second)
	t2 := time.Tick(1 * time.Second)

	go func(tick <- chan time.Time) {
		for {
			<- tick
			bank.Deposit(100)
			fmt.Printf("A %d\n", bank.Balance())
		}
	}(t1)

	for {
		<- t2
		bank.Deposit(200)
		fmt.Printf("B %d\n", bank.Balance())
	}
}
