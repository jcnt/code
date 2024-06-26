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
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {

	s := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert","Einstein","time"},
		{"Isaac","Newton","apple"},
		{"Stephen","Hawking","blackhole"},
		{"Marie","Curie","radium"},
		{"Charles","Darwin","fittest"},
	}

//	for i := range s {
//		fmt.Printf("%s\t\t", s[i][0])
//		fmt.Printf("%s\t\t", s[i][1])
//		fmt.Printf("%s\n", s[i][2])
//	}

	for i := range s {
		n := s[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 40))
		}
	}
}
