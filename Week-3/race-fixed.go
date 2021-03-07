package main

import (
	"fmt"
	"sync"
)

// Shared resource
var balance = 0

// Simple function which increments
// balance by 1
func deposit(wg *sync.WaitGroup, m *sync.Mutex) {
	// critical section start
	m.Lock()

	balance = balance + 1

	m.Unlock()
	// critical section end

	wg.Done()
}

func main() {

	// Mutex to be used
	var m sync.Mutex
	var w sync.WaitGroup

	// Spawn 1000 go routines
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go deposit(&w, &m)
	}

	w.Wait()
	fmt.Println("Bank balance is", balance)
}
