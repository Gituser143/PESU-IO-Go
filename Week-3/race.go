package main6

import (
	"fmt"
	"sync"
)

var balance = 0

// Simple function which increments
// balance by 1
func deposit(wg *sync.WaitGroup) {
	// critical section start
	balance = balance + 1
	// critical section end
	wg.Done()
}

func main() {

	var w sync.WaitGroup

	// Spawn 1000 go routines
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go deposit(&w)
	}

	w.Wait()
	fmt.Println("Bank balance is", balance)
}
