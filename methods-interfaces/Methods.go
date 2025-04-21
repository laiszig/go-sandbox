package main

import (
	"fmt"
	"math"
)

// Go does not have classes. We define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Remember: a method is just a function with a receiver argument.
// Here's Abs written as a regular function with no change in functionality.
// func Abs(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
