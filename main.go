package main

import (
	"fmt"
)

func main() {

	age := 45

	// comparators

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	// if statements
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is equal to or over 40")
	}

	// nested if statements
	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // if index == 1, it will break out of the if statement and continues back at pos 2
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break // breaks out of the loop completely, does not continue
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}