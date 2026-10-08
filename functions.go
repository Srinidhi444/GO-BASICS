package main

// this is functions in go
//When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

// func can return multiple values. Here we return two strings.
func swap(x, y string) (string, string) {
	return y, x
}

// A naked return is a return statement with no arguments. It returns the named return values. Naked returns should be used only in short functions, as they can harm readability.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
