package main

import (
	"fmt"
	"time"

	"github.com/untref-ayp2/taller-go/03-elementos-basicos/ejemplos/04-condicionales/condicionales"
)

func main() {
	var numero int
	fmt.Print("Ingrese un entero: ")
	fmt.Scanln(&numero)

	condicionales.Condicional(numero)
	condicionales.SwtichBasico()
	condicionales.SwitchSinCondicion(time.Now().Local().Hour())
	condicionales.SwitchMultiple(' ')
	condicionales.SwitchFallthrough(2)
}
