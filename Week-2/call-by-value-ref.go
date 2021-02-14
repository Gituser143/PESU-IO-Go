package main

import "fmt"

func byValue(a int) {
	a += 5
}

func byReference(a *int) {
	*a += 5
}

func main() {
	x := 10
	byValue(x)
	fmt.Println(x)

	byReference(&x)
	fmt.Println(x)
}
