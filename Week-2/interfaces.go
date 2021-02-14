package main

import (
	"fmt"
	"math"
)

// Interface decleration
type geometry interface {
	getName() string
	area() float64
	perim() float64
}

// function to work on interface
func measure(g geometry) {
	fmt.Println("Shape:", g.getName())
	fmt.Println("Area:", g.area())
	fmt.Println("Perim:", g.perim())
}

// custom type rect
type rect struct {
	width, height float64
	name          string
}

// Receivers for rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r rect) getName() string {
	return r.name
}

// custom type circle
type circle struct {
	radius float64
	name   string
}

// Receivers for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) getName() string {
	return c.name
}

func main() {
	r := rect{width: 3, height: 4, name: "Rectangle"}
	c := circle{radius: 5, name: "Circle"}

	measure(r)
	measure(c)
}
