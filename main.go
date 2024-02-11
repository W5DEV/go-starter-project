package main

import "fmt"

var someName = "hello"

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	nameOne = "peach"
	nameThree = "bowser"

	nameFour := "yoshi"

	fmt.Println(someName, nameOne, nameTwo, nameThree, nameFour)

	// ints (integers)
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 3838383838388383.99
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}