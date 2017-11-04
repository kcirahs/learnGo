package main

import "fmt"

const kmToMiles float64 = 1.6

func main() {
	var (
		km float64
		exit string
	)
	fmt.Print("Enter kilometers cycled: ")
	fmt.Scan(&km)
	miles := km / kmToMiles
	fmt.Println(km, " km ", miles, " miles.")
	fmt.Println("Type anything and press enter to exit")
	fmt.Scan(&exit)

}
