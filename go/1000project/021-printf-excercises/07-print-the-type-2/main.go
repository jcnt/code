// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
)

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func main() {
	pi := math.Pi
	fmt.Printf("Type of %.2f is %T\n", pi, pi)
}
