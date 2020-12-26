package main

import (
	"fmt"
)

func fnReceiveMessage(message string, channel chan string) {
	// Send an message in the channel
	channel <- message
}

func main() {
	// Define an new channel
	channelMessages := make(chan string)

	// go func() {channelMessages <- "Message to the channel"}()
	go fnReceiveMessage("Message to the channel", channelMessages)

	// Reading the message
	msg := <-channelMessages

	// Printing the message
	fmt.Println(msg)
}