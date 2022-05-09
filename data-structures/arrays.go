package main

import "fmt"

func main() {
	// array: inmutable and fixed size - more efficient
	// slice: mutable and variable size

	// array
	var array [4]int
	array[0] = 10
	array[1] = 11
	fmt.Println("array", (array))
	// Length and max capacity
	fmt.Println(len(array), cap(array))

	// slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)

	fmt.Println(slice[0])
	fmt.Println((slice[:2]))
	fmt.Println((slice[2:4]))
	fmt.Println((slice[4:]))

	// append
	slice = append(slice, 6)
	fmt.Println("slice append:", slice)
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println("slice append array", slice)
}
