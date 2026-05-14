package main

import "fmt"

func main() {
	s := []int{3, 7, 2, 9, 5}
	var max int
	Maximo(s, &max)
	fmt.Println(max)
}

func Maximo(numeros []int, resultado *int) {
	if len(numeros) == 0 {
		return
	}
	*resultado = numeros[0]
	for _, n := range numeros {
		if n > *resultado {
			*resultado = n
		}
	}
}
