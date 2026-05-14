package figuras

import "fmt"

type Cuadrado struct {
	posicion Punto // esquina inferior izquierda
	lado     int
}

func NewCuadrado(posicion Punto, lado int) Cuadrado {
	return Cuadrado{posicion: posicion, lado: lado}
}

func (c *Cuadrado) Mover(incX, incY int) {
	c.posicion.Mover(incX, incY)
}

func (c Cuadrado) Perimetro() int {
	return 4 * c.lado
}

func (c Cuadrado) Area() int {
	return c.lado * c.lado
}

func (c Cuadrado) String() string {
	return fmt.Sprintf("Cuadrado: [%v, %v]", c.posicion.String(), c.lado)
}
