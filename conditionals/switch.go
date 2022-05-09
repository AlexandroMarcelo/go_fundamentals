package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switch num := random(); num {
	case 0, 1, 2:
		fmt.Println(num, "-> greater or equal than 0 and lower than 3")
	case 3, 4, 5, 6:
		fmt.Println(num, "-> between 3 and 6")
		fallthrough // execute the bellow case code as well
	case 7, 8, 9:
		fmt.Println(num, "-> greater or equal than 7")
		fallthrough
	default:
		fmt.Println("Default case")
	}

	// Without conditions
	// value := 1
	switch value := 1; {
	case value > 10:
		fmt.Println(">10")
	case value < 0:
		fmt.Println("<0")
	default:
		fmt.Println("value >=0 and <= 10")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
