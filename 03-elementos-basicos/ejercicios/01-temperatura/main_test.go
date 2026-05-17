package main

import "testing"

func TestMostrarTemperatura(t *testing.T) {
	temp, cong := mostrarTemperatura()
	if temp != 36.5 {
		t.Errorf("temperatura = %f; esperado 36.5", temp)
	}
	if cong != 0 {
		t.Errorf("congelacion = %d; esperado 0", cong)
	}
}
