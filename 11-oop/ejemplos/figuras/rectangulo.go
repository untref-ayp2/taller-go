package figuras

import "fmt"

type Rectangulo struct {
	esquinaInicial, esquinaFinal Punto
}

func NewRectangulo(esquinaInicial, esquinaFinal Punto) Rectangulo {
	return Rectangulo{esquinaInicial: esquinaInicial, esquinaFinal: esquinaFinal}
}

func (r Rectangulo) Mover(incX, incY int) {
	r.esquinaInicial.Mover(incX, incY)
	r.esquinaFinal.Mover(incX, incY)
}

func (r Rectangulo) Base() int {
	return r.esquinaFinal.x - r.esquinaInicial.x
}

func (r Rectangulo) Altura() int {
	return r.esquinaFinal.y - r.esquinaInicial.y
}

func (r Rectangulo) Perimetro() int {
	return 2*r.Base() + 2*r.Altura()
}

func (r Rectangulo) Area() int {
	return r.Base() * r.Altura()
}

func (r Rectangulo) String() string {
	return fmt.Sprintf("Rectangulo: [%v, %v]", r.esquinaInicial.String(), r.esquinaFinal.String())
}
