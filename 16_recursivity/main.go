package main

import "fmt"

func fact(n int) int {
	fmt.Println("Calculando el factorial de: ", n)

	if n == 0 || n == 1 {
		fmt.Println("El factorial de: ", n, " es: 1")
		return 1
	}

	var res = n * fact (n-1)
	fmt.Println("El factorial de: ", n, " es: ", res)

	return res
}

func main() {
	fmt.Println(fact(4))
}