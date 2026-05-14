package figuras

type Rectangulo struct {
	Base   float64
	Altura float64
}

func NewRectangulo(base, altura float64) Rectangulo {
	return Rectangulo{Base: base, Altura: altura}
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
	return 2*r.Base + 2*r.Altura
}
