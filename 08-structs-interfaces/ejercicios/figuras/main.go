package main

type Punto struct {
	X, Y float64
}

func NewPunto(x, y float64) Punto {
	// TODO: implementar
	return Punto{}
}

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Rectangulo struct {
	EsqInfIzq, EsqSupDer Punto
}

func NewRectangulo(infIzq, supDer Punto) Rectangulo {
	// TODO: implementar
	return Rectangulo{}
}

func (r Rectangulo) Area() float64 {
	// TODO: implementar
	return 0
}

func (r Rectangulo) Perimetro() float64 {
	// TODO: implementar
	return 0
}

type Circulo struct {
	Centro Punto
	Radio  float64
}

func NewCirculo(centro Punto, radio float64) Circulo {
	// TODO: implementar
	return Circulo{}
}

func (c Circulo) Area() float64 {
	// TODO: implementar
	return 0
}

func (c Circulo) Perimetro() float64 {
	// TODO: implementar
	return 0
}

type Cuadrado struct {
	Rectangulo
}

func NewCuadrado(infIzq Punto, lado float64) Cuadrado {
	// TODO: implementar
	return Cuadrado{}
}

func main() {
	// Usá este espacio para probar tu implementación
}
