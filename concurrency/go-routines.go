package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup
func say(text string, wg *sync.WaitGroup) {
	// Say to the program to execute this next line until everything else has been executed
	defer wg.Done()

	fmt.Println(text)
}

func main() {
	// interact with the go routine efficient but complex to mantain it
	var wg sync.WaitGroup
	fmt.Println("Hello")

	wg.Add(1)
	go say("World!", &wg)

	wg.Wait()

	// Anonymous function
	go func(text string) {
		fmt.Println(text)
	}("Inside")

	time.Sleep(time.Second * 1) // not recommended
}
