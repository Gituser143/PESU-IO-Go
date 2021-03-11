package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, y := "10", "a"

	a, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error in converting x!")
		os.Exit(1)
	}

	b, err := strconv.Atoi(y)
	if err != nil {
		fmt.Println("Error in converting y!")
		os.Exit(1)
	}

	fmt.Println("Sum is", a+b)
}
