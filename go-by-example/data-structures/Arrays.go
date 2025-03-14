package main

import "fmt"

// Less common than Slices
func main() {

	// Zero-valued by default
	var a [5]int
	fmt.Println("emp:", a)

	// Setting a value
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len returns length
	fmt.Println("len:", len(a))

	// declare and initialize
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// compiler counts number of elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// It's possible to compose types to build multidimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Create and initialize multidimensional arrays
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
