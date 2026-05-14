package main

import "testing"

func TestExtraerOK(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	r, err := Extraer(s, 1, 3)
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	esperado := []int{20, 30}
	for i := range r {
		if r[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, r)
			break
		}
	}
}

func TestExtraerInvalido(t *testing.T) {
	s := []int{10, 20, 30}
	_, err := Extraer(s, -1, 2)
	if err == nil {
		t.Error("indice negativo deberia dar error")
	}
}
