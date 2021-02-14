package main

import "fmt"

func main() {
	// Declaring an empty slice
	// var variableName []type
	var slice []int

	fmt.Println("Slice before appends")
	fmt.Println("Slice:", slice, "Length:", len(slice))

	// Adding values to slices
	// slice = append(slice, elements)
	slice = append(slice, 1)
	slice = append(slice, 2)

	fmt.Println("Slice after appends")
	fmt.Println("Slice:", slice, "Length:", len(slice))

	// Declaring and intitialising a slice
	slice2 := []string{"Hello", "There", "General", "kenobi"}

	// Deriving slice from other slice
	// You can also derive slices from arrays
	slice3 := slice2[1:3]

	fmt.Println("Slice 2", slice2)
	fmt.Println("Slice 3", slice3)
}
