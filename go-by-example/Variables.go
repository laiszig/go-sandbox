package main

import "fmt"

// Variables are explicitly declared and used by the compiler
func main() {

	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go infers the type of initialized variables
	var d = true
	fmt.Println(d)

	// variables without initialization are zero-valued
	var e int
	fmt.Println(e)

	// := declares and initializes a variable - syntax only available inside functions
	f := "apple"
	fmt.Println(f)
}
