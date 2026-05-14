package main

import "testing"

func TestEliminarDuplicados(t *testing.T) {
	s := []int{1, 2, 2, 3, 3, 3, 4}
	s = EliminarDuplicados(s)
	esperado := []int{1, 2, 3, 4}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestEliminarDuplicadosSinDuplicados(t *testing.T) {
	s := []int{1, 2, 3}
	s = EliminarDuplicados(s)
	esperado := []int{1, 2, 3}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestEliminarDuplicadosVacio(t *testing.T) {
	s := []int{}
	s = EliminarDuplicados(s)
	if len(s) != 0 {
		t.Errorf("esperado [], obtuve %v", s)
	}
}
