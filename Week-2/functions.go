package main

import "fmt"

// Simple function with no
// parameters and no return type
func hello() {
	fmt.Println("Hello!")
}

// Function taking 2 integer
// parameters and returning 1 integer
func sum(a, b int) int {
	return a + b
}

// Function taking and returning
// 2 integer parameters
func multiplyAndSum(a, b int) (int, int) {
	return a * b, a + b
}

func main() {
	hello()

	x, y := 5, 10

	sum := sum(x, y)
	prod, _ := multiplyAndSum(x, y)

	// fmt.Println(sum, prod)
}
