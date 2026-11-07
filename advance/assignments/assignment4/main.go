package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, ch chan<- string, duration time.Duration) {
	select {
	case <-time.After(duration):
		res := "FETCHED DATA SUCCESSFULLY"

		select {
		case ch <- res:
			fmt.Println("Response sent successfully")
		case <-ctx.Done():
			fmt.Println("Request cancelled before sending")
		}

	case <-ctx.Done():
		fmt.Println("Request cancelled:", ctx.Err())
	}
}

func main() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	ch := make(chan string)

	go fetchData(ctx, ch, 5*time.Second)

	select {
	case res := <-ch:
		fmt.Println(res)

	case <-ctx.Done():
		fmt.Println("Main:", ctx.Err())
	}

	fmt.Println("Total time:", time.Since(start))
}
