package main

import "fmt"

func main() {
	total := 0
	for i := 1; i < 16; i++ {
		if i%3 == 0 {
			total += i
			fmt.Println(i, total)
		} else if i%5 == 0 {
			total += i
			fmt.Println(i, total)
		}
	}
}
