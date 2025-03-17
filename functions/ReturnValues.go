package main

import "fmt"

func main() {
	fmt.Println(split(17))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// Return values may be named - If so, they are treated as variables defined at the top of the function
// These names should be used to document the meaning of the return values
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// The function divide() has two named return values: quotient and err.
// They are assigned values inside the function
// The return statement is used without arguments, which implicitly returns the named values.
func divide(a, b float64) (quotient float64, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero is not allowed")
		return
	}
	quotient = a / b
	return // implicit return of quotient and err
}
