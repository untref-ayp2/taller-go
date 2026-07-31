package main

import (
	"os"
	"testing"
)

func TestCopiarArchivo(t *testing.T) {
	os.WriteFile("_origen.txt", []byte("contenido"), 0644)
	defer os.Remove("_origen.txt")
	defer os.Remove("_destino.txt")

	err := CopiarArchivo("_origen.txt", "_destino.txt")
	if err != nil {
		t.Fatal("error inesperado:", err)
	}

	datos, _ := os.ReadFile("_destino.txt")
	if string(datos) != "contenido" {
		t.Errorf("contenido = %q; esperado %q", string(datos), "contenido")
	}
}
