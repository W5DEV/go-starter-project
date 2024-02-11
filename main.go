package main

import "fmt"


func main() {

	// maps (like a javascript object)
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"drink": 3.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as key types
	phonebook := map[int]string{
		123456789: "mariio",
		234567891: "luigi",
		345678912: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456789])

	phonebook[123456789] = "bowser"
	phonebook[234567891] = "yoshi"

	fmt.Println(phonebook)
}