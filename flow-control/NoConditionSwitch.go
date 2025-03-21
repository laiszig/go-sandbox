package main

import (
	"fmt"
	"time"
)

// Switch without a condition is the same as switch true.
// Clean way to write long if-then-else chains.
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
