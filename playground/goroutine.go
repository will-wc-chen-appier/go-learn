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
