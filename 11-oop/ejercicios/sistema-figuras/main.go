package main

import (
	"fmt"

	"github.com/untref-ayp2/taller-go/11-oop/ejercicios/sistema-figuras/figuras"
)

func main() {
	rect := figuras.NewRectangulo(4, 5)
	cua := figuras.NewCuadrado(3)

	mostrar(rect)
	mostrar(cua)
}

func mostrar(f figuras.Figura) {
	fmt.Printf("Area: %.2f, Perimetro: %.2f\n", f.Area(), f.Perimetro())
}
