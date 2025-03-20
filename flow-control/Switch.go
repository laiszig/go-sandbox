package main

import (
	"fmt"
	"runtime"
)

// Runs first case whose value is equal to the condition expression
// Go's switch only runs the selected case, not all the cases that follow
// Break statement is provided automatically in Go.
// Switch cases need not be constants, and values involved need not be integers
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
