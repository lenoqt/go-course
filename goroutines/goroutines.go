package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

// Expanding usage of wg with mutex
func sayHelloWg() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	go sayHello()
	time.Sleep(100 * time.Microsecond)

	// Using closures
	var msg = "Hello from anon func"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye on anon func without arguments" // This will print instead "Raise condition and it is bad"
	time.Sleep(100 * time.Microsecond)
	// Use arguments instead but this still a bad practice, use wait groups instead
	var msg2 = "Hello using arguments on anon func"
	go func(msg string) {
		fmt.Println(msg)
	}(msg2)
	msg2 = "Goodbye on anon func with arguments"
	time.Sleep(100 * time.Microsecond)

	// Using waitgroups
	var msg3 = "Hello using waitgroup"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg3)
	msg3 = "Goodbye from anon func using waitgroup"
	wg.Wait()

	// Expanding wg with mutex "This is still bad"

	runtime.GOMAXPROCS(100) // This can cause problems
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHelloWg()
		m.Lock()
		go increment()
	}
	wg.Wait()

	// Viewwing threads
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
