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
	start := time.Now()

	tasks := []Task{
		{Id: 1, Duration: 2 * time.Second},
		{Id: 2, Duration: 1 * time.Second},
		{Id: 3, Duration: 3 * time.Second},
		{Id: 4, Duration: 1 * time.Second},
		{Id: 5, Duration: 2 * time.Second},
	}

	ch := make(chan *Task, len(tasks))
	wg := sync.WaitGroup{}

	for _, task := range tasks {
		wg.Add(1)
		go processTask(task, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Printf("Task %d completed\n", result.Id)
	}

	fmt.Println("Total time:", time.Since(start))
}
