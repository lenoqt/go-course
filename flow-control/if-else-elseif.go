package main

import (
	"fmt"
	"math"
)

func main() {
	if !false {
		fmt.Println("This is false")
	}
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// Initializer syntax
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	number := 50
	guess := 30
	// Logical operator
	if guess < 1 {
		fmt.Println("The guess must greater than 1")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
