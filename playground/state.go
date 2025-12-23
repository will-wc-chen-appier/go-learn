package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
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

// stateful goroutines: simulate a scenario where we use goroutine channels as locks
// setup scene: 1. Many concurrent readers 2. Some concurrent writers 3. Shared state
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func statefulGoroutineTest() {
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			// when reads and writes both have elements, go runtime chooses one randomly
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	var readCount uint64
	var writeCount uint64

	// 100 goroutines to repeatedly issue read requests
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(10),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 goroutines to repeatedly issue write requests
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(10),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readCount)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeCount)
	fmt.Println("writeOps:", writeOpsFinal)
}
