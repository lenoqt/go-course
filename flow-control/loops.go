package main

import (
	"fmt"
)

func main() {
	// Simple loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	// Bad practice
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	// Having no initializer and counter scoped to the main function
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	// while-like loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop + break/continue statement
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 10 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// Nested loop and label
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// Iterating over collections
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	person := map[string]string{
		"Name":        "Gustavo",
		"Surname":     "Barrios",
		"Nationality": "Venezuelan",
	}
	for k, v := range person {
		fmt.Println(k, v)
	}

	str := "Hello!"
	for _, v := range str {
		fmt.Println(string(v)) // Having just `v` here will print the int8 char value
	}

}
