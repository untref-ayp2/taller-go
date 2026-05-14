package main

import "fmt"

func main() {
	var cociente, resto int
	Dividir(10, 3, &cociente, &resto)
	fmt.Println(cociente, resto)
}

func Dividir(dividendo, divisor int, cociente, resto *int) {
	if divisor == 0 {
		return
	}
	*cociente = dividendo / divisor
	*resto = dividendo % divisor
}
