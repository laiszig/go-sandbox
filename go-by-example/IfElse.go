package main

import "fmt"

func main() {

	// basic condition
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// If without else
	if 8%4 == 0 {
		fmt.Println("8 division by 4")
	}

	// Logical operators || and &&
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement can precede conditionals;
	// Any variables declared in this statement are available in the current and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// No need for parenthesis around conditions, but braces are required
}
