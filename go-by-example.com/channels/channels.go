package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one
// goroutine and receive those values into another goroutine.

func main() {
	// Create a new channel
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel.
	msg := <-messages

	fmt.Println(msg)
}
