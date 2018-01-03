package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 != 0 {
			if i%5 != 0 {
				fmt.Println(i)
			} else {
				fmt.Println("Buzz")
			}
		} else if i%5 != 0 {
			fmt.Println("Fizz")

		} else {
			fmt.Println("FizzBuzz")
		}
	}
}
