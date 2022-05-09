package mypackage

import (
	"fmt"
)

// public all the names, the first case is in upper case
// private all the names, the first case is in lower case

// type carPrivate struct {
// 	brand string
// 	model int
// }

// CarPublic with public access
type CarPublic struct {
	Brand string
	Model int
	// color int // not available public
}

// PrintMessage
func PrintMessage(car string) {
	fmt.Println("Hello Car:", car)
}

// privatePrintMessage
// func privatePrintMessage(car string) {
// 	fmt.Println("Hello Car:", car)
// }
