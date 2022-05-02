package main

import (
	"fmt"
)

const shadow int16 = 27

// Enumerated constants
const (
	e = iota
	f
	g
)

const (
	errorSpecialist = iota // can be specified as _ and it will skip zero and still work, also any expression is evaluated, to offset the enum ie. io + 5 "Check next enum"
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota // Ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {

	// Naming convention is same way as regular variables
	// Typed
	// Constant have to be assigned in compile time
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Constants can be primitive types
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// Constants can be shadowed
	const shadow int = 333
	// ^ Try commenting this to check that the print statement will take the value of the outer scope for "shadow"
	fmt.Printf("%v, %T\n", shadow, shadow)

	const n1 = 42
	const n2 = 27
	fmt.Printf("%v, %T\n", n1+n2, n1+n2)

	// Enumerated constants
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)

	// if the value is not initialized to a enum member it will take value of 0 to circumvent this the first member can be set to a error value
	var specialistTypeNoInit int
	fmt.Printf("%v\n", specialistTypeNoInit == catSpecialist) // Should be false

	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)

	filesize := 4000000000.
	fmt.Printf("%.2fGB\n", filesize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("Second enum val: %b\n", isHeadquarters)
	fmt.Printf("Roles: %b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
