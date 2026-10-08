package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// this is pointers
	i, j := 1, 2
	p := &i
	fmt.Println(*p)
	*p = 2
	fmt.Println(i)
	p = &j
	*p = 3
	fmt.Println(j)
	// ------------------------
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// -----------------
	// arrays
	var a [10]string
	a[0] = "SRI"
	a[1] = "OM"
	fmt.Println(a[0], a[1])
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices
	// This selects a half-open range which includes the first element, but excludes the last one.
	var s []int = primes[1:4]
	fmt.Println(s)
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	// Other slices that share the same underlying array will see those changes.

	// slice literals is like array literals without the length.
	// This is an array literal without a length, and it is a slice.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Print(q[:])

	// ranges
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println(m["Bell Labs"])

}
