package main

import "fmt"

/*
*
Zero or more arguments
type comes AFTER the variable name
return type AFTER arguments
*/
func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func add(x int, y int) int {
	return x + y
}

// When two or more CONSECUTIVE parameters share a type - we can omit the type from all but the last
func sub(x, y int) int {
	return x - y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}
