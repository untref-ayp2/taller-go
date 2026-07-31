package main

import "testing"

func TestDividirOK(t *testing.T) {
	r, err := Dividir(10, 2)
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	if r != 5 {
		t.Errorf("Dividir(10,2) = %f; esperado 5", r)
	}
}

func TestDividirCero(t *testing.T) {
	_, err := Dividir(10, 0)
	if err == nil {
		t.Error("dividir por cero deberia devolver error")
	}
}
