package main

import "testing"

func TestSumarPunteros(t *testing.T) {
	x, y := 3, 7
	r := SumarPunteros(&x, &y)
	if r != 10 {
		t.Errorf("SumarPunteros(3,7) = %d; esperado 10", r)
	}
}

func TestSumarPunterosPrimeroNil(t *testing.T) {
	y := 7
	r := SumarPunteros(nil, &y)
	if r != 7 {
		t.Errorf("SumarPunteros(nil,7) = %d; esperado 7", r)
	}
}

func TestSumarPunterosSegundoNil(t *testing.T) {
	x := 3
	r := SumarPunteros(&x, nil)
	if r != 3 {
		t.Errorf("SumarPunteros(3,nil) = %d; esperado 3", r)
	}
}

func TestSumarPunterosAmbosNil(t *testing.T) {
	r := SumarPunteros(nil, nil)
	if r != 0 {
		t.Errorf("SumarPunteros(nil,nil) = %d; esperado 0", r)
	}
}

func TestSumarPunterosNegativos(t *testing.T) {
	x, y := -5, 5
	r := SumarPunteros(&x, &y)
	if r != 0 {
		t.Errorf("SumarPunteros(-5,5) = %d; esperado 0", r)
	}
}
