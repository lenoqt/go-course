package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {

	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	// Pointer aritmethic, to do it use "unsafe" package
	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// Using new
	var ms1 *myStruct
	ms1 = new(myStruct)
	// to work with the  struct
	ms1.foo = 42
	fmt.Println(ms1.foo)

	// With slices and maps behaviour is different since they don't actually contain the data, they contain a pointer to it. They are automatically dereferenced.
	s := []int{1, 2, 3}
	f := s
	fmt.Println(s, f)
	s[3] = 42
	fmt.Println(s, f)

	aMap := map[string]string{"foo": "bar", "baz": "buz"}
	g := aMap
	fmt.Println(aMap, g)
	aMap["foo"] = "qux"
	fmt.Println(aMap, g)

}
