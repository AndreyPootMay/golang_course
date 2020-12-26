package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Pause
func pause() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	reader.ReadString('\n')
}

func main() {
	// The ticker are the same as a timer but they have an repetitive cycle
	ticker := time.NewTicker(500 * time.Millisecond)

	// Create a channel bool to control
	go func() {
		for {
			select {
				case <-control:
					// Ending
					return
				case t := <-ticker.C:
					fmt.Println("Ticker: ", t)
					fmt.Println("Press Enter to continue...")
			}
		}
	}()

	pause()
	control<-true
	fmt.Println("Ticker stopped!")
}