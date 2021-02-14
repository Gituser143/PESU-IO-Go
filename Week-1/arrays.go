package main

import "fmt"

func main() {
	// Declaring an empty array
	// var variableName [size]type
	var arr [2]int

	// Assigning values to array elements
	arr[0] = 1
	arr[1] = 2

	fmt.Println("Array:", arr, "Length:", len(arr))

	// Declaring and intitialising an array
	arr2 := [...]string{"Hello", "There"}

	// In the above declartion ... is used to
	// infer size, you can explicitly specify
	// a number too. Eg: a := [3]int{1, 2, 3}

	fmt.Println("Array 2:", arr2, "Length:", len(arr2))
}
