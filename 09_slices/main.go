package main

import "fmt"

func main()  {
	// Así se declara un slice
	s := make([]string, 3)
	fmt.Println("Empty: ", s)
	fmt.Println("Empty:", s[0])

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Elementos de s: ", s)
	fmt.Println("Elemento de índice 2: " , s[2])
	fmt.Println("Longitud de s: ", len(s))

	s = append(s, "d");
	s = append(s, "e", "f")

	fmt.Println("Elementos después de agregar:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Slice copiado", c)

	l := s[2:5]
	fmt.Println("Elementos del 2 al 4: ", l)

	l = s[:5]
	fmt.Println("Elementos antes del 5: ", l)

	l = s[2:]
	fmt.Println("Elementos a partir del 2: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("Elementos de T: ", t)

	matriz := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		matriz[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matriz[i][j] = i + j
		}
	}

	fmt.Println("Matriz: ", matriz)
}