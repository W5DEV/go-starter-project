package main

import (
	"fmt"
	"math"
)

// Go global functions take in single variable arguments usually

func sayGreeting(n string) {
	fmt.Printf("Good morning, %v \n", n)
} 

func sayBye(n string) {
	fmt.Printf("Goodbye, %v \n", n)
} 

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
} 

// Need to specify what type of value you are returning before the function code:
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	sayGreeting("Bob")

	sayBye("Jim")


	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f and Circle 2 is %0.3f \n", a1, a2)
}