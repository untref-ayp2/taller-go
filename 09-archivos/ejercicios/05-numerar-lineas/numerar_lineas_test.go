package main

import (
	"os"
	"testing"
)

func TestNumerarLineas(t *testing.T) {
	os.WriteFile("_entrada.txt", []byte("hola\nmundo\n"), 0644)
	defer os.Remove("_entrada.txt")
	defer os.Remove("_salida.txt")

	err := NumerarLineas("_entrada.txt", "_salida.txt")
	if err != nil {
		t.Fatal("error inesperado:", err)
	}

	datos, _ := os.ReadFile("_salida.txt")
	esperado := "1: hola\n2: mundo\n"
	if string(datos) != esperado {
		t.Errorf("contenido = %q; esperado %q", string(datos), esperado)
	}
}
