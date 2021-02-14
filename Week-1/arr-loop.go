package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	// Iterating over a slice
	for index, value := range arr {
		fmt.Println(index, value)
	}

	sum := 0
	// Use _ when we don't require a variable
	for _, value := range arr {
		sum += value
	}

	fmt.Println(sum)
}
