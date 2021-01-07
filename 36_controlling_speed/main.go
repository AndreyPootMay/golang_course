package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating a channel with a buffer of 5 integers
	requests := make(chan int, 5)

	// Adding 5 numbers in the channel
	for i := 1; i <= 5; i++ {
		fmt.Println("Adding requests: ", i)
		requests <- i
	}

	// Closing the channel
	close(requests)

	// Creating a limiter
	limiter = := time.Tick(1000 * time.Millisecond)

	// Cycle for processing each response in the channel
	for xRequest := range requests {
		<-limiter
		fmt.Println("Request attended: ", xRequest, time.Now())
	}

	// Creating another channel with time values
	timeRequests := make(chan time.Time, 3)

	// Cycle to add the channel
	for i := 0; i < 3; i++ {
		request := time.Now()
		fmt.Println("Triggering request directly: ", request)
		timeRequests <- request
	}

	fmt.Println("Shooting GoRoutine...")

	// GoRoutine for proccessing the channel
	go func() {
		for t := range time.Tick(5000 * time.Millisecond) {
			fmt.Println("Shooting request at GoRoutine", t)
			timeRequests <- t
		}
	}()

	// Cycle for proccessing the responses
	for req := range xRequest {
		fmt.Println("Attending request limiter", req)
	}
}