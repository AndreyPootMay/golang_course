package main

import (
	"fmt"
	"os/exec"
)

// Function to exec a command
func execCommand(command string) string {
	var result = ""

	// Executing command
	commandResult := exec.Command("setup.exe", command)

	// Obtain the results
	_, err := commandResult.Output()

	if err != nil {
		result = err.Error()
	} else {
		// channel <- string(commandOut)
		result = "Command: " + command + " executed."
	}

	// Return
	return result
}

// Function that simulates a worker
func worker(name string, tasks, chan string, results chan string) {
	for command := range tasks {
		fmt.Println("Worker: ", name, " executing task ", command)
		results <- execCommand(command)
	}
}

// Main function
func main() {
	const tasksTotal = 5

	// channels
	tasksChannel := make(chan string, tasksTotal)
	resultsChannel := make(chan string, tasksTotal)

	// Execute three workers
	fmt.Println("Executing workers")
	go worker("Andrey", tasksChannel, resultsChannel)
	go worker("Eduardo", tasksChannel, resultsChannel)
	go worker("Maria", tasksChannel, resultsChannel)

	// Execute the tasks
	tasksChannel <- "dir c:\\windows\\System32 /s"
	tasksChannel <- "dir c:\\gradle /s"
	tasksChannel <- "dir c:\\python34 /s"
	tasksChannel <- "dir c:\\go\\course //s"

	// Closing channel
	close(tasksChannel)

	// Proccessing results
	for a := 1; a <= tasksTotal; a++ {
		fmt.Println("[", <-resultsChannel , "]")
	}

	// Closing results channel
	close(resultsChannel)

	// Output
	fmt.Println("Finished program...")
}
