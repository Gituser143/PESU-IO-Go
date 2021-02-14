package main

import "fmt"

// Creating a new type side
type side float64

// using side in a struct type
type rectangle struct {
	length side
	width  side
}

// receiver for type rectangle
func (r rectangle) getArea() float64 {
	return float64(r.length * r.width)
}

// receiver for type side
func (s side) getSquare() side {
	return s * s
}

func main() {

	var l, w side = 4, 5
	r1 := rectangle{length: l, width: w}

	area := r1.getArea()
	lSquare := l.getSquare()

	fmt.Println("Area:", area)
	fmt.Println("Square:", lSquare)
}
