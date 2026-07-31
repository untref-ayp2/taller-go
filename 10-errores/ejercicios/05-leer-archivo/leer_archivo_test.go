package main

import (
	"errors"
	"os"
	"testing"
)

func TestLeerArchivoOK(t *testing.T) {
	os.WriteFile("_test.txt", []byte("hola"), 0644)
	defer os.Remove("_test.txt")

	contenido, err := LeerArchivo("_test.txt")
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	if contenido != "hola" {
		t.Errorf("contenido = %q; esperado %q", contenido, "hola")
	}
}

func TestLeerArchivoNoExiste(t *testing.T) {
	_, err := LeerArchivo("_no_existe.txt")
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("esperado os.ErrNotExist, obtuve %v", err)
	}
}
