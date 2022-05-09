package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("Start")
	// defer fmt.Println("Middle") // Executes after main function but before main function return, Executes in LIFO
	// fmt.Println("End")

	// LIFO Demonstration
	defer fmt.Println("Start")
	defer fmt.Println("Middle")
	defer fmt.Println("End")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// defer takes the value at the time is called not at the time is executes
	a := "Start"
	defer fmt.Println(a)
	a = "Else"
}
