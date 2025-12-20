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

// create a limiter that at most accept 3 requests at the same time, effectively allowing burst
// kinda like a token-bucket behavior, where we add a token to the bucket if its vacant at a steady rate
// and each request takes a token with it if there's one to take
func burstTest() {
	limiter := make(chan time.Time, 3)
	for range 3 {
		limiter <- time.Now()
	}

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			limiter <- t
		}
	}()

	for i := 1; i <= 5; i++ {
		token := <-limiter
		fmt.Println("request", <-requests, token, time.Now())
	}
}
