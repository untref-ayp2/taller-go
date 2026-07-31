package main

import "testing"

func TestRotar(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Rotar(s, 2)
	esperado := []int{3, 4, 5, 1, 2}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestRotarCero(t *testing.T) {
	s := []int{1, 2, 3}
	Rotar(s, 0)
	esperado := []int{1, 2, 3}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}

func TestRotarCompleto(t *testing.T) {
	s := []int{1, 2, 3}
	Rotar(s, 3)
	esperado := []int{1, 2, 3}
	for i := range s {
		if s[i] != esperado[i] {
			t.Errorf("esperado %v, obtuve %v", esperado, s)
			break
		}
	}
}
