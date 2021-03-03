package main

import "fmt"

// Custom type
type car struct {
	model    string
	engine   string
	seats    int
	fuel     int
	distance float32
}

// Receivers for car
func (c car) drive() {
	for c.fuel > 0 {
		c.fuel -= 1
		c.distance += 12.5
	}
}

func (c car) service() int {
	if c.engine == "v8" {
		return 2000
	} else {
		return 1000
	}
}

func (c car) refuel(fuel int) {
	c.fuel += fuel
}

func main() {
	// Decleration of car
	volvo := car{model: "S60", engine: "v6", seats: 4, fuel: 20, distance: 0}
	fmt.Println(volvo)

	volvo.drive()
	fmt.Println(volvo)

	volvo.refuel(50)
	fmt.Println(volvo)

	cost := volvo.service()
	fmt.Println(cost, volvo)
}
