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

type mutexCounters struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *mutexCounters) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock() // always unlock on return, despite of scenarios such as panic
	c.counters[name]++
}

// mutex serve more complex cases, for example, when multiple counters are wrapped together, and can only be access one at a time
func mutexTest() {
	counter := mutexCounters{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			counter.inc(name)
		}
	}

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	wg.Wait()
	fmt.Println(counter.counters)
}
