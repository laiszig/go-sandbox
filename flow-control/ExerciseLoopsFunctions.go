package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	threshold := 1e-6

	for i := 0; i < 10; i++ {
		newZ := z - (z*z-x)/(2*z)
		fmt.Printf("Iteration %d: z = %f\n", i+1, newZ)
		if math.Abs(newZ-z) < threshold {
			break // Stop if the change is smaller than the threshold
		}
		z = newZ
	}

	return z
}

// Given a number x, we want to find the number z for which
// z^2 is most nearly x.
func main() {
	fmt.Println(Sqrt(2))
}
