package main

import "fmt"

// Creating custom type side
type side float64

// using custom type side
// in a struct type
type rectangle struct {
	length side
	width  side
}

func main() {
	r1 := rectangle{length: 4, width: 5}
	fmt.Println(r1)
}
