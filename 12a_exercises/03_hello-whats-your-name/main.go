package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Plese type your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
}
