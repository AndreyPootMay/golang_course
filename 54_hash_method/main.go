package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	c := "this is a string"

	h := sh1.New()

	h.Write([]byte(c))

	sb := h.Sum(nil)

	fmt.Println(c)
    fmt.Printf("%x\n", sb)
}