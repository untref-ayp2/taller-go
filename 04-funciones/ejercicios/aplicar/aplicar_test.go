package main

import "testing"

func TestAplicar(t *testing.T) {
	doblar := func(n int) int { return n * 2 }
	resultado := Aplicar([]int{1, 2, 3}, doblar)
	esperado := []int{2, 4, 6}
	for i := range resultado {
		if resultado[i] != esperado[i] {
			t.Errorf("Aplicar([1,2,3], doblar) = %v; esperado %v", resultado, esperado)
			break
		}
	}
}

func TestAplicarVacio(t *testing.T) {
	resultado := Aplicar([]int{}, func(n int) int { return n })
	if len(resultado) != 0 {
		t.Errorf("Aplicar([], ...) = %v; esperado []", resultado)
	}
}
