package main

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Rectangulo struct {
	Ancho, Alto float64
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
	Radio float64
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

func NuevoCuadrado(lado float64) Cuadrado {
	// TODO: implementar
	return Cuadrado{}
}

func main() {
	// Usá este espacio para probar tu implementación
}
