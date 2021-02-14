package main

import "fmt"

func main() {
	x := 10
	y := false
	z := "hi!"

	if x > 5 {
		fmt.Println("X is greater than 5")
	}

	if x == 20 {
		fmt.Println("20!")
	} else if x < 20 && !y {
		fmt.Println("< 20")
	} else {
		fmt.Println(z)
	}
}
