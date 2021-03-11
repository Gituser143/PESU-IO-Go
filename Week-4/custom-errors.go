package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("main.divide: Cannot divide by 0!")
		// Alternate
		// return 0, fmt.Errorf("Cannot divide %v by %v", a, b)
	}

	return a / b, nil
}

func main() {
	x, y := 10, 2
	z, err := divide(x, y)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quotient:", z)

	a, b := 15, 0
	c, err := divide(a, b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quotient:", c)
}
