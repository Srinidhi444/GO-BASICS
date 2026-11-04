package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d completed job %d\n", id, job)
		results <- job
	}
}
