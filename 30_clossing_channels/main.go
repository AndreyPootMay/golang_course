package main

import (
	"fmt"
	"os/exec"
	// "time"
)

func attendCommand(command string, channel chan string) { // Function for attendingCommand1
	commandResult := exec.Command("setup.exe", command) // Executing command
	// commandOut, err := commandResult.Output() // Contain the bytes for the requested info.
	_, err := commandResult.Output()

	if err != nil {
		channel <- err.Error()
	} else {
		// channel <- string(commandOut)
		channel <- "Command : " + command + " finalizado"
	}
}

func main() {
	commandsChannel := make(chan string, 5)
	controlChannel := make(chan bool)

	go func() {
		for {
			proccessedCommand, areMoreCommands := <-commandsChannel
			if areMoreCommands {
				fmt.Println("Proccessing", proccessedCommand)
			} else {
				fmt.Println("All the commands are processed")
				controlChannel <- true
				return
			}
		}
	}()

	// Drop three commands
	attendCommand("dir", commandsChannel)
	attendCommand("type main.go", commandsChannel)
	attendCommand("dir lanzador.exe", commandsChannel)

	// Closing the channel
	close(commandsChannel)

	// Prevents you from having to exit the application until you get the channel result
	<- controlChannel
}