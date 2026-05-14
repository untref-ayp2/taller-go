package main

import "testing"

func TestInicializarArreglo(t *testing.T) {
	var arr [5]int
	InicializarArreglo(&arr)
	esperado := [5]int{0, 2, 4, 6, 8}
	if arr != esperado {
		t.Errorf("esperado %v, obtuve %v", esperado, arr)
	}
}
