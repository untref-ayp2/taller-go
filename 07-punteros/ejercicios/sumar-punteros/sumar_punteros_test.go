package main

import "testing"

func TestSumarPunteros(t *testing.T) {
	x, y := 3, 7
	r := SumarPunteros(&x, &y)
	if r != 10 {
		t.Errorf("SumarPunteros(3,7) = %d; esperado 10", r)
	}
}

func TestSumarPunterosNegativos(t *testing.T) {
	x, y := -5, 5
	r := SumarPunteros(&x, &y)
	if r != 0 {
		t.Errorf("SumarPunteros(-5,5) = %d; esperado 0", r)
	}
}
