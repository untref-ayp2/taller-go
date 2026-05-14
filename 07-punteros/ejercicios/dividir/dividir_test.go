package main

import "testing"

func TestDividir(t *testing.T) {
	var cociente, resto int
	Dividir(10, 3, &cociente, &resto)
	if cociente != 3 || resto != 1 {
		t.Errorf("Dividir(10,3) = (%d,%d); esperado (3,1)", cociente, resto)
	}
}

func TestDividirExacto(t *testing.T) {
	var cociente, resto int
	Dividir(10, 2, &cociente, &resto)
	if cociente != 5 || resto != 0 {
		t.Errorf("Dividir(10,2) = (%d,%d); esperado (5,0)", cociente, resto)
	}
}
