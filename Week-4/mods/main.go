package main

import (
	"fmt"

	"github.com/Gituser143/mods/farewell"
	"github.com/Gituser143/mods/greeting"
)

func main() {
	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)

	greeting.SayHello(name)
	farewell.SayBye(name)
}
