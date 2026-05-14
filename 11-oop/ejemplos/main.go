package main

import (
	"fmt"

	"github.com/untref-ayp2/taller-go/11-oop/ejemplos/figuras"
)

func mostrar(f figuras.Figura) {
	fmt.Println(f)
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}

func main() {
	p1 := figuras.NewPunto(0, 0)
	p2 := figuras.NewPunto(10, 5)

	r := figuras.NewRectangulo(p1, p2)
	c := figuras.NewCuadrado(p1, 10)

	arreglo_figuras := [2]figuras.Figura{r, c}

	for _, f := range arreglo_figuras {
		mostrar(f)
	}
}
