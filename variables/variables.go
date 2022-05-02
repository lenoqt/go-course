package main

import (
	"fmt"
	"strconv"
)

// When declaring variables outside a function it has to be declared in the following way
var l float32 = 232.3

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane"
	doctorNumber int    = 3
	season       int    = 11
) // These variables are scoped only in this package, upper case letter that will enable to be visible on other packages

var PackageName string = "main"

var jk int = 26

func main() {
	// All variables here are only scoped to this function
	// When you want to declare in the scope but not initialize
	var i int
	i = 42
	fmt.Println(i)

	// When you want to initialize but go doesn't have enough information to inference the type of the var
	var j int = 32
	fmt.Println(j)

	// When you need to initialize the variable and the type can be inferenced
	k := 44
	fmt.Println(k)

	// Redeclaring variables
	var jk int = 42
	// jk := 13
	// Try uncommenting this ^ it will throw an error, because there's no new variable and it is declared on line 34, this is called shadowing
	jk = 32
	fmt.Println(jk)

	// kk := 11 // variables always have to be used, and it is a compile time error
	// ^ Try uncommenting this

	// Converting types *casting*
	fmt.Printf("%v, %T\n", l, l)
	i = int(l)

	// Casting float/int to string
	fmt.Printf("%v, %T\n", jk, jk)
	var kl string
	kl = string(jk) // This doesn't cast to string, strings are a alias of a stream of bytes, and the converted 32 to string, looks for what is a unicode code for that number
	fmt.Printf("%v, %T\n", kl, kl)
	// To convert to string strconv is used
	kl = strconv.Itoa(jk)
	fmt.Printf("%v, %T\n", kl, kl)

}
