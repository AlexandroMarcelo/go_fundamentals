package main

import (
	"fmt"
)

func main() {

	// Declaration inside the if (garbage collection optimized)
x	if a, b := 10, 20; a > b {
		fmt.Println("a greater than b")
	} else if a == b {
		fmt.Println("a equal b")
	} else {
		fmt.Println("a lower than b")
	}

}
