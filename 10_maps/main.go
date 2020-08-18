package main

import "fmt"

func main()  {
	// Así se crea el mapa (Diccionarios en python)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Mapa: ", m)
	fmt.Println("K1 :", m["k1"])
	v1 := m["k1"]
	fmt.Println("V1 :", v1)

	fmt.Println("Longitud del mapa: ", m)

	delete(m, "k2")

	fmt.Println("El mapa después de eliminar: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map: ", n)
}