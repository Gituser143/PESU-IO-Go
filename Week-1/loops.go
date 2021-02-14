package main

import "fmt"

func main() {
	arr := []int{}

	// 3 condition for loop
	// for <initialisation>; <condition>; <update>
	for i := 0; i < 10; i++ {
		arr = append(arr, i*i)
	}

	y := len(arr) - 1

	// While loop
	// for <condition>
	for y >= 0 {
		arr[y]++
		y--
	}

	// Infinite loop
	for {
		fmt.Println("Infinite!")
	}
}
