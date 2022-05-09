package main

import "fmt"

// private if the first lette is lower case. Public if the first case is upper case
// private structure
type car struct {
	brand string
	model int
}

type pc struct {
	brand   string
	ram     int
	storage int
}

func (_pc pc) printPcBrand() {
	fmt.Println(_pc.brand)
}

func (_pc *pc) duplicateRam() {
	_pc.ram = _pc.ram * 2
}

// Stringers: overwrite the behaviour of the print function for the struct
func (_pc pc) String() string {
	return fmt.Sprintf("PC %s, has %d GB RAM, %d GB Storage\n", _pc.brand, _pc.ram, _pc.storage)
}

func main() {
	car1 := car{brand: "BMW", model: 1998}
	fmt.Println("car1", car1)

	// Empty class
	var car2 car
	car2.brand = "Chevrolet"
	fmt.Println("car2", car2) // model is 0 because zero value (not defined value for int)

	// Working with pointers
	myPc := pc{ram: 16, storage: 256, brand: "ASUS"}
	fmt.Println("Declaration", myPc)
	myPc.printPcBrand()

	myPc.ram = 32
	myPc.duplicateRam()
	fmt.Println("After", myPc)

	// Personalization
	my2Pc := pc{ram: 8, storage: 1024, brand: "MSI"}
	fmt.Println("Declaration", my2Pc)
	// fmt.Printf("%+v", my2Pc) // also work to print in an explicit way the struct: {brand:MSI ram:8 storage:1024}

}
