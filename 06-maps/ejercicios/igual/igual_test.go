package main

import "testing"

func TestIgualIguales(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"a": 1, "b": 2}
	if !Igual(a, b) {
		t.Errorf("mapas iguales deberian ser iguales")
	}
}

func TestIgualDiferentes(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"a": 1, "b": 3}
	if Igual(a, b) {
		t.Errorf("mapas diferentes deberian ser diferentes")
	}
}

func TestIgualDistintoTamano(t *testing.T) {
	a := map[string]int{"a": 1}
	b := map[string]int{"a": 1, "b": 2}
	if Igual(a, b) {
		t.Errorf("mapas de distinto tamano deberian ser diferentes")
	}
}
