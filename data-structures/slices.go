package main

import "fmt"

func main() {
	// array: inmutable and fixed size - more efficient
	// slice: mutable and variable size

	// slice
	slice := []string{"Alex", "Raúl", "Juan"}
	fmt.Println("slice:", slice)

	// index and value
	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Println()

	// get the index only
	for i := range slice {
		fmt.Println(i)
	}

	text := "Alexandro"
	fmt.Println(text, "Reversed =", reverseText(text))
}

func reverseText(text string) string {
	var reversed string

	for i := len(text) - 1; i >= 0; i-- {
		// fmt.Println(text[i]) // ascii code of that character
		reversed += string(text[i]) // must be converted to string due character represented as ascii code
	}

	return reversed
}
