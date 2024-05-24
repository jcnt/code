// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

func main() {
	var names []string
	var distances []int
	var data []uint8
	var ratios []float64
	var alives []bool

	fmt.Printf("%-10s: %T %d %t\n", "names", names, len(names),names == nil)
	fmt.Printf("%-10s: %T %d %t\n", "distances", distances, len(distances),distances == nil)
	fmt.Printf("%-10s: %T %d %t\n", "data",data, len(data), data == nil)
	fmt.Printf("%-10s: %T %d %t\n", "ratios", ratios, len(ratios), ratios == nil)
	fmt.Printf("%-10s: %T %d %t\n", "alives",alives, len(alives), alives == nil)
}
