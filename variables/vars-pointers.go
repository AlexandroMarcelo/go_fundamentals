package main

import (
	"fmt"
)

func main() {
	name := "Alexandro"
	course := "Getting started with GO"

	fmt.Println("\nHi", name, "your current course is", course)

	updateCourse(course)
	fmt.Println("You’re currently watching", course)

	updateCoursePointer(&course)
	fmt.Println("You’re currently watching pointer", course)
}

func updateCourse(course string) string {
	course = "Getting Started with Docker"
	fmt.Println("Updated course to", course)
	return course
}

func updateCoursePointer(course *string) string {
	*course = "Getting Started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
