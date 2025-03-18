package main

import (
	"fmt"
	"math/cmplx"
)

// Just like imports, variable declarations may be "factored" into blocks
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// The int, uint, and uintptr can be 32-bit or 64-bit systems, depending on the system
// For Integer we use int unless we need a sized or unsigned integer type.
func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
