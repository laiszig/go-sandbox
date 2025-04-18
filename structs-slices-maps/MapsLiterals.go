package main

import "fmt"

// Map literals are like struct literals, but the keys are required.
type Vertex5 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex5{
	"Bell Labs": Vertex5{
		40.68433, -74.39967,
	},
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	"Google": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m2)
}
