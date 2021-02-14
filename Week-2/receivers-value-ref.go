package main

import "fmt"

type list []int

// Receiver using call by value
func (l list) append(num int) {
	l = append(l, num)
}

// Receiver using call by reference
func (l *list) appendRef(num int) {
	*l = append(*l, num)
}

func main() {
	l1 := list{1, 2, 3}

	l1.append(4)
	fmt.Println(l1)

	l1.appendRef(4)
	fmt.Println(l1)
}
