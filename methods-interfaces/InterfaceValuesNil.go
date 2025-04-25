package main

import "fmt"

// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
// In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
// Note that an interface value that holds a nil concrete value is itself non-nil.
type I3 interface {
	M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I3

	var t *T3
	i = t
	describe2(i)
	i.M()

	i = &T3{"hello"}
	describe2(i)
	i.M()
}

func describe2(i I3) {
	fmt.Printf("(%v, %T3)\n", i, i)
}
