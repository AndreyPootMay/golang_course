package main

import "fmt"

func main() {
	slice := []int{3, 4, 5, 6, 8}
	fmt.Println(slice[3]) // 6
	// fmt.Println(slice[7])
	data := slice[2]

	if data == 5 {
		panic("OMG is the channel 5, what's wrong?!")
	} else {
		fmt.Println(data)
	}

	fmt.Println("End")
}