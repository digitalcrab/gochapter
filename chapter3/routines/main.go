package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")

	go func() {
		fmt.Println("World")
	}()

	// wait for a goroutine
	time.Sleep(time.Second)
}
