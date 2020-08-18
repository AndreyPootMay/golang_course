package main

import "fmt"

if 10%2 == 0 {
	fmt.Println("10 es par")
} else {
	fmt.Println("10 es impar")
}

if 8%4 == 0 {
	fmt.Println("8 es divisible entre 4")
}

if num := 9, num < 0 {
	fmt.Println("es negativo")
} else if num < 10 {
	fmt.Println("tiene un digito")
} else {
	fmt.Println("es mayor que 9")
}