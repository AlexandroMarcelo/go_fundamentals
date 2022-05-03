package main

import (
	"fmt"
)

func main() {
	randomNumber := getLowerNumber(1, 3, 7, 11)
	fmt.Println(randomNumber)
}

// Unknown number of params
func getLowerNumber(numbers ...int) int {
	lower := numbers[0]
	for _, i := range numbers {
		if i < lower {
			lower = i
		}
	}
	return lower
}
