package main

import (
	"fmt"
	"time"
)

func fetchData(ch chan<- string, duration time.Duration) {
	time.Sleep(duration)
	res := "FETCHED DATA SUCCESSFULLY"
	ch <- res
}

func main() {
	start := time.Now()
	ch := make(chan string, 1)
	go fetchData(ch, 3*time.Second)
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("TIMEOUT")
	}
	fmt.Println("total time:", time.Since(start))
}
