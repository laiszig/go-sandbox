package main

import "fmt"

// The range form of the for loop
// iterates over a slice or map
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	// When ranging over a slice, two values are returned for each iteration
	// First is the index, second is a copy of the element at that index
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
