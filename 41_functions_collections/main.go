package main

import (
	"fmt"
	"strings"
)

func getIndex(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func isIncluded(vs []string, t string) bool {
	return getIndex(vs, t) >= 0
}

func getAny(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func getAll(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if lf(v) {
			return false
		}
	}
	return true
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	var strs = []string{"peach", "apple", "grape", "plum", "pear", "orange"}

	// Index
	fmt.Println("Index of pear", Index(strs, "pear"))
	fmt.Println("Index of orange", Index(strs, "orange"))

	// Include
	fmt.Println("Is included grape: ", isIncluded(strs, "grape"))
	fmt.Println("Is included plum: ", isIncluded(strs, "plum"))

	// Returns only if have a p letter in his names
	fmt.Println("Only if they have a p: ", getAny(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// Return only if contains
	fmt.Println("Only if they contain a pe: ", getAny(strs, func(v string) bool {
		return strings.Contains(v, "pe")
	}))

	// Mapping
	fmt.Println(Map(strs, strings.ToUpper))
}