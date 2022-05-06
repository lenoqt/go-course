package main

import (
	"fmt"
)

func main() {

	// slices are dynamic size and are a projection of A
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	B := A[:]   // Slice of all elements
	C := A[3:]  // Slice from 4th element
	D := A[:6]  // Slice first 6 elements
	E := A[3:6] // Slice the 4th, 5th and 6th element
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Printf("Size and capacity of A: %v %v\n", len(A), cap(A))
	fmt.Printf("Size and capacity of B: %v %v\n", len(B), cap(B))
	fmt.Printf("Size and capacity of C: %v %v\n", len(C), cap(C))
	fmt.Printf("Size and capacity of D: %v %v\n", len(D), len(D))
	fmt.Printf("Size and capacity of E: %v %v\n", len(E), cap(E))

	// Using make
	array := make([]int, 3) // We can initialize the capacity a := make([]int, 10, 100) slice with len 10 and cap 100
	fmt.Println(array)
	fmt.Printf("Length: %v\n", len(array))
	fmt.Printf("Capacity: %v\n", cap(array))

	// Increasing capacity
	big := []int{}
	fmt.Println(big)
	fmt.Printf("Length: %v\n", len(big))
	fmt.Printf("Capacity: %v\n", cap(big))
	big = append(big, 1)
	fmt.Println(big)
	// Using variadic function
	fmt.Printf("Length: %v\n", len(big))
	fmt.Printf("Capacity: %v\n", cap(big))
	big = append(big, 2, 3, 4, 5)
	fmt.Println(big)
	fmt.Printf("Length: %v\n", len(big))
	fmt.Printf("Capacity: %v\n", cap(big))
	// Using spread out slice like *array in Python
	big = append(big, []int{6, 7, 8, 9}...)
	fmt.Println(big)
	fmt.Printf("Length: %v\n", len(big))
	fmt.Printf("Capacity: %v\n", cap(big))

	// Removing an arbitraty value of array
	fmt.Println(big)
	bigNoMiddle := append(big[:4], big[5:]...)
	fmt.Println(bigNoMiddle)
	fmt.Printf("Length: %v\n", len(bigNoMiddle))
	fmt.Printf("Capacity: %v\n", cap(bigNoMiddle))
	fmt.Printf("Original array again which shows UB: %v\n", big)
}
