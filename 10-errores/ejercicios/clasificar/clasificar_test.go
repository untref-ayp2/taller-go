package main

import "testing"

func TestClasificarNino(t *testing.T) {
	c, err := ClasificarEdad(5)
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	if c != "nino" {
		t.Errorf("edad 5 -> %q; esperado nino", c)
	}
}

func TestClasificarNegativo(t *testing.T) {
	_, err := ClasificarEdad(-1)
	if err == nil {
		t.Error("edad negativa deberia devolver error")
	}
}
