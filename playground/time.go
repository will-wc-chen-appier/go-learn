package main

import (
	"fmt"
	"time"
)

func timerTest() {
	timer := time.NewTimer(time.Second)

	go func() {
		<-timer.C
		fmt.Println("message received")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("timer stopped")
	}

	// reset := timer.Reset(time.Second)
	// if reset {
	// 	fmt.Println("timer reset")
	// }

	time.Sleep(2 * time.Second)
}

func tickerTest() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at: ", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
