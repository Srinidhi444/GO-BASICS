package main

import "fmt"

func whileloop() {

	// go doesnt have a while loop
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
	switchdemo()
}
