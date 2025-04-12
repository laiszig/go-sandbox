package main

import "fmt"

// Append: 2 parameters
// T and T values to append to the slice
// Result: slice containing all the elements.
// If the backing array is too small to fit all values,
// A bigger array is allocated, the return value will point
// to the newly allocated array.
func main() {

	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
