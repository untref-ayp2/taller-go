package main

import (
	"math"
	"testing"
)

func TestRectanguloArea(t *testing.T) {
	r := Rectangulo{Ancho: 3, Alto: 4}
	if r.Area() != 12 {
		t.Errorf("Area = %f; esperado 12", r.Area())
	}
}

func TestRectanguloPerimetro(t *testing.T) {
	r := Rectangulo{Ancho: 3, Alto: 4}
	if r.Perimetro() != 14 {
		t.Errorf("Perimetro = %f; esperado 14", r.Perimetro())
	}
}

func TestCirculoArea(t *testing.T) {
	c := Circulo{Radio: 5}
	esperado := math.Pi * 25
	if c.Area() != esperado {
		t.Errorf("Area = %f; esperado %f", c.Area(), esperado)
	}
}

func TestCirculoPerimetro(t *testing.T) {
	c := Circulo{Radio: 5}
	esperado := 2 * math.Pi * 5
	if c.Perimetro() != esperado {
		t.Errorf("Perimetro = %f; esperado %f", c.Perimetro(), esperado)
	}
}

func TestCuadradoArea(t *testing.T) {
	c := NuevoCuadrado(4)
	if c.Area() != 16 {
		t.Errorf("Area = %f; esperado 16", c.Area())
	}
}

func TestCuadradoPerimetro(t *testing.T) {
	c := NuevoCuadrado(4)
	if c.Perimetro() != 16 {
		t.Errorf("Perimetro = %f; esperado 16", c.Perimetro())
	}
}
