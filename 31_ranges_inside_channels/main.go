package main

import "fmt"

func main() {
	// Creating a channel
	queue := make(chan string, 2)
	// Keep the 2 data bytes in the channel
	queue <- "one"
	queue <- "two"

	// Close the cannel
	close(queue)

	// Using an range with "for" to obtain the data.

	for elem := range queue {
		// print the obtained data
		fmt.Println(elem)
	}
}