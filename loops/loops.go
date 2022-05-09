package main

import "fmt"

func main() {
	// There are only for loop which can represent a for loop and while loop
	// For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")

	// While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("")

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
		// continue
		if i == 2 {
			fmt.Println("Is 2. continue")
			continue
		}

		// break
		if i == 8 {
			fmt.Println("Is 8. break")
			break
		}
	}

}
