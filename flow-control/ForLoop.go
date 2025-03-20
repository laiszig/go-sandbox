package main

import "fmt"

// init statement: executed before the first iteration
// condition expression: evaluated before every iteration
// post statement: executed at the end of every iteration
func main() {

	sum := 0
	// No parentheses surrounding the three components
	// Braces are always required
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	// The init and post statements are optional.
	// we can even drop the semicolons in this case
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// if we omit the condition - infinite loop
	for {

	}
}
