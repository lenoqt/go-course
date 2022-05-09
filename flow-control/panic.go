package main

import (
	"fmt"
	"log"
	// "net/http"
)

func deferAfter() {
	// deferred statements still executes.
	fmt.Println("Start")
	defer fmt.Println("This was deferred")
	panic("Something bad happened")
	fmt.Println("End")
}

func recoverDefer() {
	fmt.Println("Start")
	// Anonymous function
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() // Invoke function
	panic("Something bad happened")
	fmt.Println("End")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err) // to re-panic the function
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking")
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// deferAfter()
	// recoverDefer()
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}
