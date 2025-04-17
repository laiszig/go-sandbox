package main

import "fmt"

// A map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
type Vertex4 struct {
	Lat, Long float64
}

var m map[string]Vertex4

func main() {
	// The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Vertex4)
	m["Bell Labs"] = Vertex4{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
