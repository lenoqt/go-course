package main

import (
	"fmt"
)

func main() {
	// arrays are fixed size and determined at compile time
	grades := [...]int{95, 52, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[2] = "Mai"
	students[1] = "Arnold"
	fmt.Printf("Students: %v\n", students[1])
	fmt.Printf("Size of students: %v\n", len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// Arrays are considered values
	// here will create a copy
	a := [...]int{1, 2, 3}
	b := a // Change this to &a so it points to array a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

}
