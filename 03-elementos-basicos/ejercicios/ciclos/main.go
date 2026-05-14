package main

import "fmt"

func main() {
	numeros := []int{3, 7, 2, 9, 5}
	fmt.Println("Máximo:", maximo(numeros))
}

func maximo(numeros []int) int {
	if len(numeros) == 0 {
		return 0
	}
	max := numeros[0]
	for _, n := range numeros {
		if n > max {
			max = n
		}
	}
	return max
}
