package main

import "fmt"

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
func main() {

	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
