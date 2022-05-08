package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// Embedding (Composition)
// Tags
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[0])

	// Anonymous struct
	aNewDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aNewDoctor)
	anotherDoctor := aNewDoctor // This will create a copy, in order to manipulate the original object, a reference to it must be passed.
	anotherDoctor.name = "Tom Barker"
	fmt.Println(aNewDoctor)
	fmt.Println(anotherDoctor)
	// manipulate without creating copies
	sameDoctor := &aNewDoctor
	sameDoctor.name = "Gustavo Barrios"
	fmt.Println(aNewDoctor)
	fmt.Println(sameDoctor)

	// Embedding
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println(b)

	// Declaring as literal syntax
	bird := Bird{
		Animal: Animal{Name: "Ju", Origin: "Australia"},
		Speed:  48,
		CanFly: false,
	}
	fmt.Println(bird)
	// Reflect of a struct tag
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
