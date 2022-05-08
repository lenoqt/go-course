package main

import (
	"fmt"
)

func main() {
	// Cases can't be overlapped (repeated )
	switch 2 {
	case 1, 5, 10:
		fmt.Println("One, five or ten")
	case 2, 4, 6:
		fmt.Println("Two, four or six")
	default:
		fmt.Println("Another number")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("One, five or ten")
	case 2, 4, 6:
		fmt.Println("Two, four or six")
	default:
		fmt.Println("Another number")
	}

	// Using tag, here cases can be overlapped
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
		fallthrough // This is logicallyless
	case i <= 20:
		fmt.Println("Less or equal to twenty")
	default:
		fmt.Println("Greater than twenty")
	}

	// Typed switch
	var j interface{} = [3]int{}
	switch j.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}

	var k interface{} = 1
	switch k.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This is unreachable")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")

	}
}
