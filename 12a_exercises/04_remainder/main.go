package main

import "fmt"

func main() {
	var (
		big   int
		small int
	)
	fmt.Print("Please enter a number: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a larger number: ")
	fmt.Scan(&big)
	if big > small {
		fmt.Println(big, "/", small, " has a remainder of ", big%small)
	} else {
		fmt.Println(big, " is not greater than ", small)
	}
}
