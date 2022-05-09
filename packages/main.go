package main

import (
	"fmt"

	pkCar "amarcelo.com/ps-go/packages/mypackage"
)

func main() {
	var myCar pkCar.CarPublic
	myCar.Brand = "BMW"
	myCar.Model = 1998

	fmt.Println(myCar)
	pkCar.PrintMessage("BMW")
}
