package main

import "fmt"

func generate(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending", i)
		ch <- i
	}
	close(ch)
}

func main() {
	chan1 := make(chan int, 0)

	go generate(chan1)

	num := <-chan1
	fmt.Println("Got", num)

	for num := range chan1 {
		fmt.Println("Got", num)
	}
}
