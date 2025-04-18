package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// If v or ok have not yet been declared you could use this short declaration form.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	// If key is in m, ok is true, if not ok is false
	// If key is not in the map, then v is the zero value for the map's element type
}
