package main

import "testing"

func TestInvertir(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Invertir(s)
	esperado := []int{5, 4, 3, 2, 1}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestInvertirVacio(t *testing.T) {
	s := []int{}
	Invertir(s)
}
