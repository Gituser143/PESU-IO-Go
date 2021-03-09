package main

import (
	"fmt"
	"time"
)

func writeEmail() {
	fmt.Println("Finished writing an email!")
}

func orderFood() {
	fmt.Println("Finished ordering food!")
}

func bookTickets() {
	fmt.Println("Finished booking tickets!")
}

func main() {
	go writeEmail()
	go orderFood()
	go bookTickets()
	fmt.Println("Started routines!")
	time.Sleep(10 * time.Second)
}
