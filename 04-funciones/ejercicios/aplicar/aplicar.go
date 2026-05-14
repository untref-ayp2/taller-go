package main

import "fmt"

func main() {
	doblar := func(n int) int { return n * 2 }
	fmt.Println(Aplicar([]int{1, 2, 3}, doblar))
}

func Aplicar(numeros []int, f func(int) int) []int {
	resultado := make([]int, len(numeros))
	for i, n := range numeros {
		resultado[i] = f(n)
	}
	return resultado
}
