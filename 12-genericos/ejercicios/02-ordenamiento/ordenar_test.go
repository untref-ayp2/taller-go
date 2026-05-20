package main

import "testing"

func TestOrdenarEnteros(t *testing.T) {
	arr := []int{4, 2, 7, 1, 9}
	OrdenarSeleccion(arr, func(a, b int) bool { return a < b })
	esperado := []int{1, 2, 4, 7, 9}
	if len(arr) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(arr))
	}
	for i, v := range arr {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestOrdenarInverso(t *testing.T) {
	arr := []int{4, 2, 7, 1, 9}
	OrdenarSeleccion(arr, func(a, b int) bool { return a > b })
	esperado := []int{9, 7, 4, 2, 1}
	if len(arr) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(arr))
	}
	for i, v := range arr {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestOrdenarStrings(t *testing.T) {
	arr := []string{"Pepe", "Ana", "Luis"}
	OrdenarSeleccion(arr, func(a, b string) bool { return a < b })
	esperado := []string{"Ana", "Luis", "Pepe"}
	if len(arr) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(arr))
	}
	for i, v := range arr {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %s, obtenido %s", i, esperado[i], v)
		}
	}
}

func TestOrdenarYaOrdenado(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	OrdenarSeleccion(arr, func(a, b int) bool { return a < b })
	esperado := []int{1, 2, 3, 4, 5}
	if len(arr) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(arr))
	}
	for i, v := range arr {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %d, obtenido %d", i, esperado[i], v)
		}
	}
}

func TestOrdenarUnElemento(t *testing.T) {
	arr := []int{42}
	OrdenarSeleccion(arr, func(a, b int) bool { return a < b })
	if arr[0] != 42 {
		t.Error("Un solo elemento debe quedar igual")
	}
}

func TestOrdenarVacio(t *testing.T) {
	arr := []int{}
	OrdenarSeleccion(arr, func(a, b int) bool { return a < b })
	// No debe panic
}

func TestOrdenarFlotantes(t *testing.T) {
	arr := []float64{3.5, 1.2, 4.8, 2.1}
	OrdenarSeleccion(arr, func(a, b float64) bool { return a < b })
	esperado := []float64{1.2, 2.1, 3.5, 4.8}
	if len(arr) != len(esperado) {
		t.Fatalf("Longitud esperada %d, obtenida %d", len(esperado), len(arr))
	}
	for i, v := range arr {
		if v != esperado[i] {
			t.Errorf("En posicion %d: esperado %f, obtenido %f", i, esperado[i], v)
		}
	}
}
