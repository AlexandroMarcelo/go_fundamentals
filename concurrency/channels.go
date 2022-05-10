package main

import "fmt"

// chan<- input
// <-chan output

func say(text string, c chan<- string) {
	c <- text // assaign the text to that channel - input
	// text = <- c output
}

func message(message string, c chan string) {
	c <- message
}

func main() {
	// Channels
	c := make(chan string, 1) // last arg how many data will be handled in the go-routine (best practices)

	fmt.Println("Hello")
	go say("World!", c)

	fmt.Println(<-c)
	fmt.Println("Bye")

	c2 := make(chan string, 3)
	c2 <- "Message1"
	c2 <- "Message2"

	// len: how many go-routines are in the channel
	// cap: max capacity in the channel
	fmt.Println(len(c2), cap(c2))

	///////////////

	// close: the channel will close for the  go runtime and will not be abel to receive more data. Good practice
	close(c2)
	// c2 <- "Message3" // will not work due channel closed

	///////////////
	// range: loop over each messages/data in the channel
	for message := range c2 {
		fmt.Println(message)
	}

	///////////////

	// select: to handle the response of each channel
	email1 := make(chan string)
	email2 := make(chan string)
	go message("email 1", email1)
	go message("email 2", email2)

	// 2 because there are two channels
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email1 message:", m1)

		case m2 := <-email2:
			fmt.Println("Email2 message:", m2)
		}
	}
}
