package main

import (
	"fmt"
	"strings"
)

func main() {
	conteo := ContarPalabras("hola mundo hola")
	fmt.Println(conteo)
}

func ContarPalabras(texto string) map[string]int {
	conteo := make(map[string]int)
	for _, p := range strings.Fields(texto) {
		conteo[p]++
	}
	return conteo
}
