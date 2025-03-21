package main

import "fmt"

// Why use pointers
// Efficient memory usage – Instead of copying large data structures, you can pass pointers, reducing memory overhead.
// Modify values inside functions – Functions in Go pass arguments by value (copying them). Using pointers allows you to modify the original data.
// Work with data structures – Pointers are essential for linked lists, trees, and other dynamic data structures.

// Function that modifies the original variable using a pointer
func double(n *int) {
	*n = *n * 2 // deferencing the pointer to modify the value
}

// You can't swap two numbers in Go using a function unless you use pointers.
// Without pointers, Go would copy the values, and the swap wouldn't persist outside the function.
func swap(a, b *int) {
	*a, *b = *b, *a // Swap values using pointers
}

func main() {

	x := 10
	fmt.Println("Before:", x)

	double(&x)               // Pass the address of x
	fmt.Println("After:", x) // Now x is modified

	x, y := 5, 10
	fmt.Println("Before:", x, y)

	swap(&x, &y)                // Passing memory addresses
	fmt.Println("After:", x, y) // Values are swapped
}
