package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// every second
	t := time.Tick(1 * time.Second)

	// every 3 seconds
	t2 := time.Tick(3 * time.Second)

	// infinitive loop
	for {
		// select allows you read from the channels (depend which one has a value now)
		select {
		case <-t:
			fmt.Println("Tick")
		case <-t2:
			os.Exit(0)
		}
	}
}
