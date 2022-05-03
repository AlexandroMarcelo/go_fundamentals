package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name   = "Alexandro"
	course string
	module = 5
	clip   int
	a      = "1"
	b      = 2
)

func main() {
	lastName := "Marcelo"
	// boolean := false // -> must be in use otherwise error while compiling
	fmt.Println("Name of the user", name, lastName, "and the course is", course)
	fmt.Println("Module and clip are set to", module, "and", clip)
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	iA, err := strconv.Atoi(a)
	if err == nil {
		total := iA + b
		fmt.Println("The total is", total)
	}

	fmt.Println("Lastname memory address is", &lastName)
	var ptr *string = &lastName // pointer variable to point to the lastName variable
	fmt.Println("Pointing lastName at address", ptr, "which holds this value", *ptr)

}
