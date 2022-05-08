package main

import (
	"fmt"
)

func main() {
	// Maps key items must be able to be tested for equality
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

	// Non-equality example
	// m := map[[]int]string{}

	// Using make
	person := make(map[string]string)
	person = map[string]string{
		"First Name":  "Gustavo",
		"Last Name":   "Barrios",
		"Nationality": "Venezuelan",
	}
	fmt.Println(person)

	// Pulling values
	fmt.Println(statePopulations["Ohio"])
	// Adding values
	statePopulations["Georgia"] = 10310371
	// The order of items on a maps are not guaranteed
	fmt.Println(statePopulations)
	// Deleting
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	// Asking for a value that doesn't exists returns 0
	fmt.Println(statePopulations["Georgia"])
	// A way to circumvent this is to use the ok notation, ok will return  false or true
	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)
	//  _, ok := statePopulations["California"]
	// fmt.Println(ok)

	// Passing maps is passed by reference
	sp := statePopulations
	delete(sp, "Texas")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
