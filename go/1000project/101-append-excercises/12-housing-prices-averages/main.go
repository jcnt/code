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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	d := strings.Split(data, "\n")

	for _, v := range strings.Split(header, separator) {
		fmt.Printf("%-15s", v)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 65))

	for _, v := range d {
		w := strings.Split(v, separator)
		for _, x := range w {
			fmt.Printf("%-15s", x)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 65))

	for _,v := range d {
		fmt.Printf("%v, %t", v[1], v[1])
	}
}
