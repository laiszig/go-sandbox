package main

import "fmt"

// Slices can be created with the make function
// That's how we create dynamically-sized arrays.
func main() {

	// Make function allocates a zeroed array
	// And returns a slice that refers to that array
	a := make([]int, 5)
	printSlice2("a", a)

	// To specify capacity, pass a third argument to make
	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
