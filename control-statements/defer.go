package main

import "fmt"

func defering() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
