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

func selectTest() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("done")
}

func timeoutTest() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

func nonBlockingTest() {
	messages := make(chan string, 1)
	signals := make(chan string, 1)

	select {
	case message := <-messages:
		fmt.Println("received message:", message)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent: ", msg)
	default:
		fmt.Println("no message sent")
	}

	signals <- "hehe"

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}

func closeTest() {
	//create a goroutine that waits for jobs to be sent over thru queue
	jobs := make(chan int, 5)
	done := make(chan string)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- "done"
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	<-done //for synchronization

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}

func rangeOverTest() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
