package main

import "fmt"

type car struct {
	model    string
	engine   string
	seats    int
	fuel     int
	distance float32
}

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

type bike struct {
	model    string
	engine   string
	fuel     int
	distance float32
}

func (b bike) drive() {
	for b.fuel > 0 {
		b.fuel -= 1
		b.distance += 12.5
	}
}

func (b bike) service() int {
	if b.engine == "b1" {
		return 500
	} else {
		return 200
	}
}

func (b bike) refuel(fuel int) {
	b.fuel += fuel
}

type vehicle interface {
	drive()
	service() int
	refuel(int)
}

func serviceAll(vehicles []vehicle) int {
	cost := 0
	for _, v := range vehicles {
		cost += v.service()
	}

	return cost
}

func main() {
	volvo := car{model: "S60", engine: "v6", seats: 4, fuel: 20, distance: 0}
	lambo := car{"Countach", "v8", 2, 10, 0}

	kawasaki := bike{model: "Ninja", engine: "b1"}
	yamaha := bike{model: "R1", engine: "b3", fuel: 5}

	myVehicles := []vehicle{volvo, lambo, kawasaki, yamaha}

	totalCost := serviceAll(myVehicles)

	fmt.Println("Total service cost:", totalCost)

}
