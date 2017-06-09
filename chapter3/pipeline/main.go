package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// the first goroutine adds the values to the naturals channel
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// the second one reads from the naturals and sends the new value to the squares
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
