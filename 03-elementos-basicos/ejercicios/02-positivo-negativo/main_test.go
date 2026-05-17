package main

import "testing"

func TestPositivo(t *testing.T) {
	r := clasificarNumero(5)
	if r != "positivo" {
		t.Errorf("clasificarNumero(5) = %q; esperado \"positivo\"", r)
	}
}

func TestNegativo(t *testing.T) {
	r := clasificarNumero(-3)
	if r != "negativo" {
		t.Errorf("clasificarNumero(-3) = %q; esperado \"negativo\"", r)
	}
}

func TestCero(t *testing.T) {
	r := clasificarNumero(0)
	if r != "cero" {
		t.Errorf("clasificarNumero(0) = %q; esperado \"cero\"", r)
	}
}
