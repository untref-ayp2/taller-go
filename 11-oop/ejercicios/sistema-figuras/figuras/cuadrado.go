package figuras

type Cuadrado struct {
	Lado float64
}

func NewCuadrado(lado float64) Cuadrado {
	return Cuadrado{Lado: lado}
}

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (c Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}
