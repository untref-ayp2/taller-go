package main

import "testing"

func TestSaludar(t *testing.T) {
	resultado := saludar("Martín")
	esperado := "¡Hola, Martín!"
	if resultado != esperado {
		t.Errorf("saludar(\"Martín\") = %q; esperado %q", resultado, esperado)
	}
}

func TestSaludarVacio(t *testing.T) {
	resultado := saludar("")
	esperado := "¡Hola, !"
	if resultado != esperado {
		t.Errorf("saludar(\"\") = %q; esperado %q", resultado, esperado)
	}
}
