package main

import (
	"fmt"
	"os/exec"
	"sync"
)

// Function to exec a command
func execCommand(command string) string {
	var result = ""

	// Executing command
	commandResult := exec.Command("setup.exe", command)

	// Obtain the results
	outputData, err := commandResult.Output()

	if err != nil {
		result = err.Error()
	} else {
		// channel <- string(commandOut)
		result = "Command: " + command + " executed -->" + string(outputData)
	}

	// Return
	return result
}

func worker(name string, task string, wg *sync.WaitGroup) {
	// Make sure about the wg closing
	defer wg.Done()

	// Message
	fmt.Println("Worker %s starting with ->%s \n", name, task)

	// Executing command
	result := execCommand(task)

	fmt.Println("Worker %s ended with ->%s \n", name, resultsChannel)
}

func main() {
	// Wg variable
	var wg sync.WaitGroup

	// Workers array
	workers := [5]string{"Andrey", "Eduardo", "Maria", "Josefa"}
	commands := [5]string{"dir c:\\windows\\system32", "type main.go", "dir a", "dir b", "dir c"}

	// Cycle for execute

	for i := 0; i < len(workers); i++ {
		wg.Add(1)
		go worker(workers[i], commands[i], &wg)
	}

	// Waiting for closing
	wg.Wait()
	fmt.Println("End of program")
}