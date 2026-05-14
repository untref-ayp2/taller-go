package main

import "testing"

func TestMaximo(t *testing.T) {
	var max int
	Maximo([]int{3, 7, 2, 9, 5}, &max)
	if max != 9 {
		t.Errorf("maximo = %d; esperado 9", max)
	}
}

func TestMaximoNegativos(t *testing.T) {
	var max int
	Maximo([]int{-5, -2, -8, -1}, &max)
	if max != -1 {
		t.Errorf("maximo = %d; esperado -1", max)
	}
}
