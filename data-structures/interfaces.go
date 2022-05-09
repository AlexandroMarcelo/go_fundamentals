package main

import "fmt"

// Interface to admit structs as parameters and implement the method it has
type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (c square) area() float64 {
	return c.base * c.base
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func calculateArea(f figure2D) {
	fmt.Println("Area", f.area())
}

func main() {
	mySquare := square{base: 2.5}
	fmt.Printf("Square: %+v\n", mySquare)
	// mySquare.area()
	calculateArea(mySquare)

	myRectangle := rectangle{base: 5, height: 2}
	fmt.Printf("Rectangle: %+v\n", myRectangle)
	// myRectangle.area()
	calculateArea(myRectangle)

	fmt.Println()

	// List of interfaces
	listInterfaces := []interface{}{"Hello", 1, 2.3, false}
	fmt.Println(listInterfaces...)
}
