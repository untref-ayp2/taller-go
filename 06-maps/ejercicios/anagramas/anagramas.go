package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SonAnagramas("listen", "silent"))
	fmt.Println(SonAnagramas("hola", "mundo"))
}

func SonAnagramas(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if len(a) != len(b) {
		return false
	}
	frecuencias := make(map[rune]int)
	for _, c := range a {
		frecuencias[c]++
	}
	for _, c := range b {
		frecuencias[c]--
		if frecuencias[c] < 0 {
			return false
		}
	}
	return true
}
