package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} 

	return initials[0], "_"

}

func main() {

	fname1, sname1 := getInitials("gilderoy lockhart")
	fmt.Println(fname1, sname1)

	fname2, sname2 := getInitials("Severus snape")
	fmt.Println(fname2, sname2)

	fname3, sname3 := getInitials("voldemort")
	fmt.Println(fname3, sname3)

}