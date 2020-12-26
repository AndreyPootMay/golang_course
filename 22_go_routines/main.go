package main

import (
	"fmt"
	"time"
)

func fnGoFunction(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(50* time.Millisecond)
		fmt.Println(s, i);
	}
}

func main() {
	go fnGoFunction("Running like a thread")
	fnGoFunction("Running normally")
}