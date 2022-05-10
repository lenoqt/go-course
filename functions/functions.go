package main

import (
	"fmt"
)

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessageIndex(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// Pass by value
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

// Pass by pointer
func sayGreetingPointer(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)

}

// Variadic params received as slice
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// Using syntax sugar
func sumReturn(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func sumReturnPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// Multiple return and idiomatic error handling
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// Methods
type greeter struct {
	greeting string
	name     string
}

// Value receiver
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// With pointer
func (g *greeter) greetPointer() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	sayMessage("Hello!")
	for i := 0; i < 5; i++ {
		sayMessageIndex("Hello!", i)
	}
	// Pass by value
	fmt.Println("Pass by value")
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(greeting, name)
	fmt.Println(name)
	// Pass by pointer
	fmt.Println("Pass by pointer")
	sayGreetingPointer(&greeting, &name)
	fmt.Println(name)

	// Variadic func
	fmt.Println("Variadic function")
	sum("The sum is", 1, 2, 3, 4, 5)
	s := sumReturn(1, 2, 3, 4, 5)
	fmt.Println("The sum using return is:", s)

	sP := sumReturnPointer(1, 2, 3, 4, 5)
	fmt.Println("The sum using pointers: ", *sP)

	// Multiple returns
	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Anonymous function best practice for async
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println("From Anonymous func ", i)
		}(i)
	}
	// func as variable
	var dividevar func(float64, float64) (float64, error)
	dividevar = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	dd, err := dividevar(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dd)

	// Methods
	fmt.Println("From methods")
	g := greeter{
		greeting: "Hello",
		name:     "Go"}
	g.greet()
	g.greetPointer()
}
