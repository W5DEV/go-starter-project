package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// strings package

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello")) // returns true/false if variable contains given string
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // replaces text in the given string (but doesn't alter the string)
	fmt.Println(strings.ToUpper(greeting)) // replaces string to uppercase 
	fmt.Println(strings.Index(greeting, "lo")) // finds the position of the given string
	fmt.Println(strings.Split(greeting, "e")) // splits the string by the given element


	// sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // sorts and alters the given slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) // finds the position as index of the given value
	fmt.Println(index)

	index2 := sort.SearchInts(ages, 80)
	fmt.Println(index2)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"} 

	sort.Strings(names) // sorts and alters the given slice
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) // gives position as index of given value
}