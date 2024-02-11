package main

import (
	"fmt"
)

func main() {

	x := 0

	// for is used for loops, even "while" loops:
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	// traditional for loops:
	for i := 0; i < 6; i++ {
		fmt.Println(x == i) // comparator
	}
	
	
	names := []string{"mario", "luigi", "yoshi", "peach"}

	for j := 0; j < len(names); j++ {
		fmt.Println(names)
	}

	// for...in example
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	// if you don't want index or value, you need to use a _ to ignore it
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string" // This does not update the value of the slice - it is merely a copy of the value of the slice position
	}

	fmt.Println(names);
}