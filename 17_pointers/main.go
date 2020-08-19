package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int)  {
	*iptr = 5
}

func main() {
	i := 10

	fmt.Println("Valor inicial: ", i)
	zeroval(i)
	fmt.Println("Valor después de zeroval: ", i)
	zeroptr(&i)
	fmt.Println("Valor después de zeroptr: ", i)

	fmt.Println("Puntero: ", &i)

	// Declaring a variable pointing to an int
	var pInt *int
	pInt = &i
	fmt.Println("pInt: ", pInt)
	fmt.Println("pInt: ", *pInt)
}