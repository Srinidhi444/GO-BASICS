package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id       int
	Duration time.Duration
}
type Result struct {
	TaskID int
}

func processTask(task Task, ch chan<- *Task, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(task.Duration)
	ch <- &task
}

func main() {
	// start := time.Now()

	// tasks := []Task{
	// 	{Id: 1, Duration: 2 * time.Second},
	// 	{Id: 2, Duration: 1 * time.Second},
	// 	{Id: 3, Duration: 3 * time.Second},
	// 	{Id: 4, Duration: 1 * time.Second},
	// 	{Id: 5, Duration: 2 * time.Second},
	// }

	// ch := make(chan *Task, len(tasks))
	// wg := sync.WaitGroup{}

	// for _, task := range tasks {
	// 	wg.Add(1)
	// 	go processTask(task, ch, &wg)
	// }

	// wg.Wait()
	// close(ch)

	// for result := range ch {
	// 	fmt.Printf("Task %d completed\n", result.Id)
	// }

	// fmt.Println("Total time:", time.Since(start))

	start := time.Now()

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send 10 jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Close results after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect completed jobs
	for result := range results {
		fmt.Printf("Received result for job %d\n", result)
	}

	fmt.Println("Total time:", time.Since(start))
}
