package main

import "fmt"

// var statement declares a list of variables; type is last
var c, python, java bool

// var declaration can include initializers, one per variable
var i, j int = 1, 2

// can be at package or function level
func main() {
	var z int
	fmt.Println(z, c, python, java)

	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var ruby, golang, csharp = true, false, "no!"
	fmt.Println(i, j, ruby, golang, csharp)
}
