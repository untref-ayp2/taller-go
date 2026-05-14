package main

import "testing"

func TestEliminar(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = Eliminar(s, 2)
	esperado := []int{1, 2, 4, 5}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestEliminarPrimero(t *testing.T) {
	s := []int{1, 2, 3}
	s = Eliminar(s, 0)
	esperado := []int{2, 3}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestEliminarFueraDeRango(t *testing.T) {
	s := []int{1, 2, 3}
	s = Eliminar(s, 10)
	esperado := []int{1, 2, 3}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}
