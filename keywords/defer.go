package main

import "fmt"

func main() {
	// defer: Make the code line to execute at the end of the program
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

}
