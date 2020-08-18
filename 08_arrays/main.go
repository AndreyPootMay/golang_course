package main

import "fmt"

func main() {
	var a[5]int
	fmt.Println("Elementos de a: ", a)
	// a = append(a, 5)

	a[4] = 100
	fmt.Println("Elementos modificados: ", a)
	fmt.Println("Elemento en 4", a[4])

	fmt.Println("Longitud del arreglo: ", len(a))

	b := [5]int{1,2,3,4,5}

	fmt.Println("Elementos de b: ", b)

	var matrix [2][3]int
	
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}

	fmt.Println("Matriz: ", matrix)

	fmt.Println("Primer renglón de la matriz: ", matrix[0])
}