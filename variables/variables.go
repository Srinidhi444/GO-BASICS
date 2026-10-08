package main

import "fmt"

var c, name, age string = "c", "John", "30"

// basic types in go
// bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr byte rune float32 float64 complex64 complex128
func main() {
	var i int = 10

	// short declarations
	k := 2

	fmt.Println(i, c, name, age, k)

	var a int
	fmt.Println("uninitialized is ->", a)

	// type conversions
	var f float64 = 3.14
	var i2 int = int(f)
	fmt.Println("converted int is ->", i2)

	// constants in go
	const Pi = 3.14
	const World = "世界"
	fmt.Println("constants are ->", Pi, World)
}
