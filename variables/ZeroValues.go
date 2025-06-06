package main

import "fmt"

/*
*
Variables declared without an explicit initial value are given their zero value.
- 0 for numeric types,
- false for booleans
- "" (the empty string) for strings.
*/
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
