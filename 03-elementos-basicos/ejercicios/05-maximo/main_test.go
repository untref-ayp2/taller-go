package main

import "testing"

func TestMaximo(t *testing.T) {
	r := maximo([]int{3, 7, 2, 9, 5})
	if r != 9 {
		t.Errorf("maximo([3,7,2,9,5]) = %d; esperado 9", r)
	}
}

func TestMaximoNegativos(t *testing.T) {
	r := maximo([]int{-5, -2, -8, -1})
	if r != -1 {
		t.Errorf("maximo([-5,-2,-8,-1]) = %d; esperado -1", r)
	}
}
