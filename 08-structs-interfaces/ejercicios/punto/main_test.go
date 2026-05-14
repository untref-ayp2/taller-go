package main

import "testing"

func TestNuevoPunto(t *testing.T) {
	p := NuevoPunto(3, 4)
	if p.X != 3 || p.Y != 4 {
		t.Errorf("NuevoPunto(3,4) = (%f,%f); esperado (3,4)", p.X, p.Y)
	}
}

func TestDistanciaOrigen(t *testing.T) {
	p := NuevoPunto(3, 4)
	d := p.DistanciaOrigen()
	if d != 25 {
		t.Errorf("DistanciaOrigen(3,4) = %f; esperado 25", d)
	}
}
