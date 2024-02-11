package main

import "fmt"

var someName = "hello"

func main() {

	// arrays
	//var ages [3]int = [3]int{20, 25, 30} // arrays must be fixed lengths and the length can never be changed

	var ages = [3]int{20, 25, 30} // shorthand

	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	names[1] = "luigi"
	
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (more like typical arrays from elsewhere) - arrays that can be manipulated

	var scores = []int{100, 50, 60}

	scores[1] = 75

	scores = append(scores, 85) // append new value to slice (cannot do this for arrays)

	fmt.Println(scores, len(scores))

	// slice ranges
	range1 := names[1:3] // new slice made of names array position 1-3 (does not include position 3)

	fmt.Println(range1)

	range2 := names[2:] // selects everything from the 2nd position to the end of the array

	fmt.Println(range2)

	range3 := names[:3] // selects everything up to but not including position 3
	
	fmt.Println(range3)

	range1 = append(range1, "koopa")

	fmt.Println(range1)

}