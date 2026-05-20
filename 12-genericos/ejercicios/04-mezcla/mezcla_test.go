package main

import "testing"

func TestMergeEnteros(t *testing.T) {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	result := MergeSlices(a, b, func(x, y int) bool { return x < y })
	esperado := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestMergePrimerVacio(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}
	result := MergeSlices(a, b, func(x, y int) bool { return x < y })
	esperado := []int{1, 2, 3}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestMergeSegundoVacio(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{}
	result := MergeSlices(a, b, func(x, y int) bool { return x < y })
	esperado := []int{1, 2, 3}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestMergeAmbosVacios(t *testing.T) {
	result := MergeSlices([]int{}, []int{}, func(x, y int) bool { return x < y })
	if len(result) != 0 {
		t.Error("Merge vacio debe devolver slice vacio")
	}
}

func TestMergeDesiguales(t *testing.T) {
	a := []int{1, 5, 9}
	b := []int{2, 3, 7, 10, 12}
	result := MergeSlices(a, b, func(x, y int) bool { return x < y })
	esperado := []int{1, 2, 3, 5, 7, 9, 10, 12}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestMergeStrings(t *testing.T) {
	a := []string{"Ana", "Luis"}
	b := []string{"Beatriz", "Pepe"}
	result := MergeSlices(a, b, func(x, y string) bool { return x < y })
	esperado := []string{"Ana", "Beatriz", "Luis", "Pepe"}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %s, obtenido %s", i, esperado[i], v)
		}
	}
}
