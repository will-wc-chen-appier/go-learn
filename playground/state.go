package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// create 50 goroutine, each adding 1000 times to ops
func atomicCounterTest() {
	var ops atomic.Uint64
	var ops2 uint64
	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
				ops2++
			}
		})
	}

	wg.Wait()

	fmt.Println("ops: ", ops.Load())
	fmt.Println("ops2:", ops2) //obvious case of race condition, interleaving of goroutine happened
}
