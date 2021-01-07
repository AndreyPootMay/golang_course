package main

import "fmt"

func main() {
	// Declaring the var ageOld
	var ageOld int

	// Request age
	fmt.Println("Type your years old")

	// Read the variable
	fmt.Scanln(&ageOld)

	// At ends, always show this message
	defer fmt.Println("Thanks, come back")

	if ageOld < 18 {
		fmt.Println("You cannot access to the secret message")
		return
	}

	fmt.Println("Hello I am a secret message that shows only if you has more than 18 years old")
}