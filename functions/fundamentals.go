package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "Alexandro"
	course := "Go Fundamentals"
	fmt.Println(mixer(author, course))
}

func mixer(author string, course string) (a string, c string) {
	author = strings.ToUpper(author)
	course = strings.ToLower(course)

	return author, course
}
