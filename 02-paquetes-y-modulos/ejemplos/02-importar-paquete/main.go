package main

import (
	"fmt"
	mat "math"
	"strings"
)

func main() {
	texto := "hola mundo"
	mayusculas := strings.ToUpper(texto)
	fmt.Println(mayusculas)

	raiz := mat.Sqrt(144)
	fmt.Printf("Raíz cuadrada de 144: %.0f\n", raiz)
}
