package main

import "fmt"

// chan<- input
// <-chan output

func say(text string, c chan<- string) {
	c <- text // assaign the text to that channel - input
	// text = <- c output
}

func main() {
	// Channels
	c := make(chan string, 1) // last arg how many data will be handled in the go-routine (best practices)

	fmt.Println("Hello")
	go say("World!", c)

	fmt.Println(<-c)
}
