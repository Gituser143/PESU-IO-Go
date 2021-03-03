package main

import "fmt"

// Variadic function, takes one
// or more values of same type
func vSum(a ...int) int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}

func main() {
	// Calling vSum with aribitrary
	// number of arguments
	s1 := vSum(1, 2, 15, 7, 10, 20)

	arr := []int{1, 5, 6, 3, 14}

	// Providing a slice as argument
	s2 := vSum(arr...)

	fmt.Println("Sum S1:", s1)
	fmt.Println("Sum S2:", s2)
}
