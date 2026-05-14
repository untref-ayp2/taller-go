package main

import (
	"testing"

	"github.com/untref-ayp2/taller-go/11-oop/ejercicios/sistema-figuras/figuras"
)

func TestRectanguloArea(t *testing.T) {
	r := figuras.NewRectangulo(4, 5)
	if r.Area() != 20 {
		t.Errorf("Area = %f; esperado 20", r.Area())
	}
}

func TestCuadradoArea(t *testing.T) {
	c := figuras.NewCuadrado(3)
	if c.Area() != 9 {
		t.Errorf("Area = %f; esperado 9", c.Area())
	}
}

func TestFigurasPolimorfismo(t *testing.T) {
	var f figuras.Figura
	f = figuras.NewRectangulo(4, 5)
	if f.Perimetro() != 18 {
		t.Errorf("Perimetro = %f; esperado 18", f.Perimetro())
	}
	f = figuras.NewCuadrado(3)
	if f.Perimetro() != 12 {
		t.Errorf("Perimetro = %f; esperado 12", f.Perimetro())
	}
}
