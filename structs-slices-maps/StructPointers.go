package main

import "fmt"

// Struct fields can be accessed through a struct pointer.
type Vertex2 struct {
	X int
	Y int
}

func main() {
	v := Vertex2{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
