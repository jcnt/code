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
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 2 {
		fmt.Println("I need one argument!")
	}

	g, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
	}

	run := g

	if run <= 10 {
		run = 10
	}

	for i := 1; i <= g; i++ {
		r := rand.Intn(run + 1)
		fmt.Printf("iteration %d with random %d\n", i, r)

		if r == g {
			fmt.Println("You won!")
			return
		}
	}
	fmt.Println("You lose!")
}
