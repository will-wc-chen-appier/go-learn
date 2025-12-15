package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(w int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", w, "started job ", j)
		time.Sleep(2 * time.Second)
		fmt.Println("worker ", w, "finished job ", j)
		results <- j
	}
}

func workerTest() {
	const numJobs int = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// fmt.Println("result: ", <-results, "received")
		<-results
	}
}

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroupTest() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker2(i)
		})
	}
	wg.Wait()
}
