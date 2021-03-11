package main3

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x, y := "10", "a"

	a, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}

	b, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is", a+b)
}
