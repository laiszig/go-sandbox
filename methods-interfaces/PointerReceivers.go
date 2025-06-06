package main

import (
	"fmt"
	"math"
)

/**
Pointer receivers are a way of associating methods with a type such
that the method receives a pointer to the type as its receiver, rather than a copy of the value.
*/

// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points.
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
