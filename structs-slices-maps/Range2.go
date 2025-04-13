package main

import "fmt"

func main() {

	// You can skip the index or value by assigning to _
	// If we only want the index, we can omit the second variable
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // ==2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
