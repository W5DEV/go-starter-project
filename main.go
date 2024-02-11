package main

import "fmt"

var someName = "hello"

func main() {

	age := 35
	name := "Bob"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("Hello, world!")
	fmt.Println("Goodbye, world!")

	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name) // grabs the variables IN ORDER
	fmt.Printf("My age is %q and my name is %q \n", age, name) // puts quotes around strings (doesn't work on variables)
	fmt.Printf("My age is %v and my name is %v \n", age, name) 
	fmt.Printf("Age is of type %T \n", age) // grabs the type of the variable
	fmt.Printf("You scored %0.1f points! \n", 255.55) // grabs the float variable provided to the set decimal places
	
	// Sprintf (save formatted strings)
	var savedString = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", savedString)
	
}