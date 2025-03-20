package main

import (
	"fmt"
	"math"
)

// Doesn't need to be surrounded by parenthesis
// Braces are required
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
