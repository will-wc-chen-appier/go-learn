package main

import (
	"fmt"
	"time"
)

// using a time.Tick to steadily allow a request at a defined pace
func steadyTest() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for i := 1; i <= 5; i++ {
		x := <-limiter
		fmt.Println("request", <-requests, x)
	}
}
