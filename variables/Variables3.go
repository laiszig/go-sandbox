package main

import "fmt"

/*
*
If a variable doesn't specify an explicit type, it is inferred from the value on the right hand side.
*/
func main() {

	var i int
	j := i // When the right hand side of the declaration is typed, the new variable is of that same type.

	// When the right hand side contains an untyped numeric constant,
	// The new variable may be an int, float64, or complex128 depending on the precision of the constan
	h := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Println(i, j, f, g, h)

}
