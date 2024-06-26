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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

func main() {

	rand.Seed(time.Now().UnixNano())

	var n int

	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	} else if len(os.Args) == 3 {
		n, _ = strconv.Atoi(os.Args[2])
	}

	for i := 0; i <= n; i++ {

		r := rand.Intn(n + 1)

		if os.Args[1] == "-v" {
			if r == n {
				fmt.Println("You won!")
				return
			} else {
				fmt.Printf("%d ", r)
			}
		} else if r == n {
			fmt.Println("You won!")
			return
		}



	}
	fmt.Println("you lose")

}
