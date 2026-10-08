package main

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// func fibonacci() func() int {
//     a, b := 0, 1

//     return func() int {
//         result := a
//         a, b = b, a+b
//         return result
//     }
// }
