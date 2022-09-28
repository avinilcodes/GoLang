package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) { //jobs is sender channel and results is receiver channel
	for j := range jobs {
		fmt.Println("Worker ", id, "Started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "Finished ", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		r := <-results
		fmt.Println(r)
	}

}

// Irrespective of jobs sequence we want to get the task/job/computation done in background
// we use worker pool.
