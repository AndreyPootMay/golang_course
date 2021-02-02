package main

import "fmt"

func add2nums(a int, b int) int {
	return a + b
}

func add3nums(a, b, c int) int {
	return a + b + c
}

func main() {
	res := add2nums(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = add3nums(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}