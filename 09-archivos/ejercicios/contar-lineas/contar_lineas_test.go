package main

import (
	"os"
	"testing"
)

func TestContarLineas(t *testing.T) {
	contenido := []byte("linea1\nlinea2\nlinea3\n")
	os.WriteFile("_test.txt", contenido, 0644)
	defer os.Remove("_test.txt")

	r := ContarLineas("_test.txt")
	if r != 3 {
		t.Errorf("ContarLineas = %d; esperado 3", r)
	}
}

func TestContarLineasVacio(t *testing.T) {
	os.WriteFile("_test_vacio.txt", []byte{}, 0644)
	defer os.Remove("_test_vacio.txt")

	r := ContarLineas("_test_vacio.txt")
	if r != 0 {
		t.Errorf("ContarLineas vacio = %d; esperado 0", r)
	}
}
