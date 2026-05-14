package main

import (
	"fmt"

	"github.com/untref-ayp2/taller-go/11-oop/ejemplos/figuras"
)

func main() {
	p := figuras.NewPunto(0, 0)
	fmt.Println(p)
	p.Mover(10, -10)
	fmt.Println(p)

	p1 := figuras.NewPunto(0, 0)
	p2 := figuras.NewPunto(5, 3)
	r := figuras.NewRectangulo(p1, p2)

	fmt.Println("Area del Rectangulo: ", r.Area())
	r.Mover(12, -10)
	fmt.Println("Area del Rectangulo: ", r.Area())
	fmt.Println("Perimetro del Rectangulo: ", r.Perimetro())
	fmt.Println(r)

	c := figuras.NewCuadrado(p1, 5)
	fmt.Println("Area del Cuadrado: ", c.Area())
	fmt.Println("Perimetro del Cuadrado: ", c.Perimetro())
	fmt.Println(c)
}
