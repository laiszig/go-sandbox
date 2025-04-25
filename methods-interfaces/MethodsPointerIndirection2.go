package main

import (
	"fmt"
	"math"
)

/*
*
Functions that take a value argument must take a value of that specific type:
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!

While methods with value receivers take either a value or a pointer as the receiver when they are called:
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/
type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex4{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex4{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
