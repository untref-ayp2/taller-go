package main

import (
	"fmt"

	"github.com/untref-ayp2/taller-go/04-funciones/ejemplos/genericas"
	"github.com/untref-ayp2/taller-go/04-funciones/ejemplos/matematicas"
)

func main() {
	genericas.Saludar("Ana")

	saludo := genericas.Saludar2("Juan")
	fmt.Println(saludo)

	x := "Mundo"
	y := "Hola"
	fmt.Println("x:", x, "| y:", y)
	x, y = genericas.Swap(x, y)
	e1, e2 := 5, 10
	e2, e1 = e1, e2

	fmt.Println(e1, e2)
	fmt.Println("Swap => x:", x, "| y:", y)

	resultado, resto := matematicas.Dividir(10, 3)
	fmt.Printf("Division: %d. Resto: %d\n", resultado, resto)

	total := matematicas.Sumar(10, 10, 10, 10, 10, 5)
	fmt.Println("El total de la suma es:", total)
}
