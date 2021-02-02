package main

import (
	"strconv"
	"fmt"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

	//`64` requiere que el resultado quepa en 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
		
	u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

	k, _ := strconv.Atoi("135")
    fmt.Println(k)

	z, e := strconv.Atoi("wat")
    fmt.Println(e)
    fmt.Println(z)
}