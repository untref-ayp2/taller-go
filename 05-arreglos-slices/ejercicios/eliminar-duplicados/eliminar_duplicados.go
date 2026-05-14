package main

import "fmt"

func main() {
	s := []int{1, 2, 2, 3, 3, 3, 4}
	s = EliminarDuplicados(s)
	fmt.Println(s)
}

func EliminarDuplicados(s []int) []int {
	if len(s) == 0 {
		return s
	}
	resultado := []int{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			resultado = append(resultado, s[i])
		}
	}
	return resultado
}
