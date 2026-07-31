package main

import "testing"

func TestPromedio(t *testing.T) {
	resultado := Promedio(4.0, 5.0, 6.0)
	esperado := 5.0
	if resultado != esperado {
		t.Errorf("Promedio([4,5,6]) = %.2f; esperado %.2f", resultado, esperado)
	}
}

func TestPromedioVacio(t *testing.T) {
	resultado := Promedio()
	if resultado != 0 {
		t.Errorf("Promedio([]) = %.2f; esperado 0", resultado)
	}
}
