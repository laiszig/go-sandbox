package main

import (
	"fmt"
	"slices"
)

// Typed only by the elements they contain - not the number of elements
func main() {

	// Uninitialized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	// create an empty slice ith non-zero length
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// set and get
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// length of the slice
	fmt.Println("len: ", len(s))

	// more operations than arrays

	// Append returns a slice containing one or more new values
	// We need to accept a return value from append - as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// Slices can be copied.
	c := make([]string, len(s))
	copy(c, s) // copy into c from s
	fmt.Println("cpy:", c)

	// slice operator - slice[low:high]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// up to but excluding :idx
	l = s[:5]
	fmt.Println("sl2:", l)

	// up from and including idx:
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize a variable for slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The slices package contains several utility functions
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Multidimensional slices.
	// The length of the inner slices can vary, unlike with arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
