package main

import "fmt"

func main() {
	greatest := max(23, 1, 0, 23, 25, 67, 129, 200, 1000)
	fmt.Println(greatest)
}

func max(numbers ...int) int{
	//fmt.Printf("%T", numbers)
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}