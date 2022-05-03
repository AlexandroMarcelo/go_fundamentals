package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Environ())
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	var username string = os.Getenv("USER")
	// username := os.Getenv("USER") // Also works for declaration

	fmt.Println("\nUsername in the OS logged in", username)
}
