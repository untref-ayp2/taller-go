package main

import "fmt"

func main() {
	fmt.Println(clasificarNumero(5))
	fmt.Println(clasificarNumero(-3))
	fmt.Println(clasificarNumero(0))
}

func clasificarNumero(n int) string {
	if n > 0 {
		return "positivo"
	} else if n < 0 {
		return "negativo"
	}
	return "cero"
}
