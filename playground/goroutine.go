package main

import (
	"fmt"
	"time"
)

func printNumbers(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func goroutineTest() {
	printNumbers("direct")

	go printNumbers("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println(time.Second)
	time.Sleep(time.Second)
	fmt.Println("done")
}

func channelTest() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}

func bufferTest() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// messagesBlocked := make(chan string)
	// messagesBlocked <- "blocked" // fatal error: all goroutines are asleep - deadlock!

	// messages := make(chan string)
	// messages <- "ping" // fatal error: all goroutines are asleep - deadlock!
	// go func() {
	// 	msg := <-messages
	// 	fmt.Println(msg)
	// }()
	// time.Sleep(time.Millisecond)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func directionTest() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
