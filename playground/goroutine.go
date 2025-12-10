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
