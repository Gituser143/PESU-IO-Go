package main

import "fmt"

func main() {
	// Decleration followed by intitialisation
	var x int
	x = 0

	// Decleration and intitialisation
	var y float64 = 1.2

	// Type inferring
	z := "Hi!"

	// Unused variable, not allowed
	a := 123

	fmt.Println("Int:", x)
	fmt.Println("Float:", y)
	fmt.Println("String:", z)
}
