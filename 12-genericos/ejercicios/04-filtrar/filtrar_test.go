package main

import "testing"

func TestFiltrarPares(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	result := Filtrar(arr, func(n int) bool { return n%2 == 0 })
	esperado := []int{2, 4, 6}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestFiltrarNada(t *testing.T) {
	arr := []int{1, 3, 5}
	result := Filtrar(arr, func(n int) bool { return n%2 == 0 })
	if len(result) != 0 {
		t.Error("Filtrar impares debe devolver slice vacio")
	}
}

func TestFiltrarTodo(t *testing.T) {
	arr := []int{2, 4, 6}
	result := Filtrar(arr, func(n int) bool { return n%2 == 0 })
	esperado := []int{2, 4, 6}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestFiltrarStrings(t *testing.T) {
	arr := []string{"Ana", "Luis", "Pepe", "Alberto", "Maria"}
	result := Filtrar(arr, func(s string) bool { return len(s) > 3 })
	esperado := []string{"Luis", "Pepe", "Alberto", "Maria"}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %s, obtenido %s", i, esperado[i], v)
		}
	}
}

func TestFiltrarVacio(t *testing.T) {
	result := Filtrar([]int{}, func(n int) bool { return true })
	if len(result) != 0 {
		t.Error("Filtrar slice vacio debe devolver slice vacio")
	}
}

func TestFiltrarConservaOrden(t *testing.T) {
	arr := []int{10, 3, 8, 1, 6, 2}
	result := Filtrar(arr, func(n int) bool { return n < 5 })
	esperado := []int{3, 1, 2}
	if len(result) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(result))
	}
	for i, v := range result {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}
