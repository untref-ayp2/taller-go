package punteros

import (
	"fmt"
)

func Incrementar(num *int) {
	*num++

	fmt.Println("La dirección de memoria del numero es: ", num)
}

func Incrementar2(num int) {
	num++
}
