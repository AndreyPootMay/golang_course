package main

import (
	"fmt"
	"os/exec"
)

func attendCommand(command string, channel chan string) { // Function for attendingCommand1
	commandResult := exec.Command("setup.exe", command) // Executing command
	commandOut, err := commandResult.Output() // Contain the bytes for the requested info.
	// _, err := commandResult.Output()

	if err != nil {
		channel <- err.Error()
	} else {
		channel <- string(commandOut)
		// channel <- "Command : " + commandOut + " finalizado"
	}
}

func main() {
	// Three channels
	var channel1 chan string = make(chan string)
	var channel2 chan string = make(chan string)
	var channel3 chan string = make(chan string)

	// Drop the thread
	go attendCommand("dir c:\\Windows\\System32 /s", channel1)
	go attendCommand("Type main.go", channel2)
	go attendCommand("dir c:\\go\\course /s", channel3)

	// Proccess the response from the three channels
	for i := 0; i < 3 ; i++ {
		select {
			case responseChannel1 := <- channel1:
				fmt.Println("Response from channel 1: ", responseChannel1)
			case responseChannel2 := <- channel2:
				fmt.Println("Response from channel 2: ", responseChannel2)
			case responseChannel3 := <- channel3:
				fmt.Println("Response from channel 3: ", responseChannel3)
		}
	}
	fmt.Println("End...")
}