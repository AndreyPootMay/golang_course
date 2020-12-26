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
	// Creating a timer
	timer1 := time.NewTimer(2 * time.Second)
	// Wait for a response
	<-timer1.C
	fmt.Println("Timer 1 executed.")

	// Creating an timer
	timer2 := time.NewTimer(2 * time.Second)

	// Calling a goRoutine
	go func() {
		// Waiting the triggering
		<-timer2.c
		fmt.Println("Timer 2 executed.")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		// Veryfying timer 2 stopping
		fmt.Println("Timer 2 stopped")
	}

	pause()
}