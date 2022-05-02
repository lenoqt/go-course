package main

import (
	"fmt"
)

func main() {

	// Booleans
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	k := 1 == 1
	l := 2 == 1

	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", l, l)

	// Every time you initialize a variable in go, it is initialized to zero
	var i bool
	fmt.Printf("%v, %T\n", i, i)

	// Numeric types

	// Integers
	// Signed integers int8 int16 int32 int64...
	sNumber := 42
	fmt.Printf("%v, %T\n", sNumber, sNumber)
	// Unsigned integers byte(uint8) uint16 uint32 uint64..
	var uNumber uint16 = 44
	fmt.Printf("%v, %T\n", uNumber, uNumber)

	// Operations (Must be same type)
	a := 10
	b := 3
	fmt.Println(a + b) // Addition
	fmt.Println(a - b) // Substraction
	fmt.Println(a * b) // Multiplication
	fmt.Println(a / b) // Division
	fmt.Println(a % b) // Modulus
	// a = 1010 ; b = 0011
	fmt.Println(a & b)  // AND     0010
	fmt.Println(a | b)  // OR      1011
	fmt.Println(a ^ b)  // EOR     1001
	fmt.Println(a &^ b) // ANDNOT 0100

	// Bit shifting
	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0

	// Floats
	// var f float32
	// ^ Try uncommenting this, it should overfloat on line 56
	f := 3.14
	f = 13.6e72
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)

	// Operations
	aa := 10.2
	bb := 3.7
	fmt.Println(aa + bb) // Addition
	fmt.Println(aa - bb) // Substraction
	fmt.Println(aa * bb) // Multiplication
	fmt.Println(aa / bb) // Division

	// Complex
	var complex complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", complex, complex)
	fmt.Printf("%v, %T\n", real(complex), real(complex))
	fmt.Printf("%v, %T\n", imag(complex), imag(complex))
	var bigC complex128 = complex128(complex)
	fmt.Printf("%v, %T\n", bigC, bigC)

	// Operations
	ac := 10 + 2i
	bc := 3 + 23i
	fmt.Println(ac + bc) // Addition
	fmt.Println(ac - bc) // Substraction
	fmt.Println(ac * bc) // Multiplication
	fmt.Println(ac / bc) // Division

	// Strings
	s := "A string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2]) // Strings are aliases for bytes and are immutable
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// String concatenation
	s1 := "Hello"
	s2 := " you!"
	fmt.Println(s1 + s2)

	// Strings can be collected to slice of bytes
	sBytes := []byte(s)
	fmt.Printf("%v, %T\n", sBytes, sBytes)

	// Rune
	r := 'a'
	fmt.Printf("%v, %T\n", r, r) // Rune are true type alias to int32 / utf-32

}
