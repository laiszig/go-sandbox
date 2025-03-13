package main

import "fmt"

func main() {

	// Most basic type - single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial/condition after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// "Do this N times" iteration over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// For without a condition will loop repeatedly until you break out of the loop or return from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// Continue to the next iteration of the loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
