package main

import (
	"os"
	"testing"
)

func TestAgregarLinea(t *testing.T) {
	os.WriteFile("_bitacora.txt", []byte("linea1\n"), 0644)
	defer os.Remove("_bitacora.txt")

	err := AgregarLinea("_bitacora.txt", "linea2")
	if err != nil {
		t.Fatal("error inesperado:", err)
	}

	datos, _ := os.ReadFile("_bitacora.txt")
	esperado := "linea1\nlinea2\n"
	if string(datos) != esperado {
		t.Errorf("contenido = %q; esperado %q", string(datos), esperado)
	}
}
