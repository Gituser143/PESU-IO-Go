package main

import (
	"fmt"
	"sync"
)

func writeEmail(wg *sync.WaitGroup) {
	fmt.Println("Finished writing an email!")
	wg.Done()
}

func orderFood(wg *sync.WaitGroup) {
	fmt.Println("Finished ordering food!")
	wg.Done()
}

func bookTickets(wg *sync.WaitGroup) {
	fmt.Println("Finished booking tickets!")
	wg.Done()
}

func main() {
	// Create WaitGroup variable
	var wg sync.WaitGroup

	// Initialise WaitGroup
	wg.Add(2)

	// Start routines
	go writeEmail(&wg)
	go orderFood(&wg)
	go bookTickets(&wg)

	fmt.Println("Started routines!")

	// Wait for routines to finish
	wg.Wait()
}
