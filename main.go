package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)

	fmt.Println(add(1, 2))
	// := creates a variable with the type of the value on the right hand side
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
