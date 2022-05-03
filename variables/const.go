package main

import (
	"fmt"
)

func main() {
	const a = 12345
	fmt.Println("The const is", a)

	// a = 54321 // -> ERROR :(
	// fmt.Println("Does not print")
}
