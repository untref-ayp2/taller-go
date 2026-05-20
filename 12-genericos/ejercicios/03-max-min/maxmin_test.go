package main

import "testing"

func TestMaximoEnteros(t *testing.T) {
	arr := []int{10, 5, 8, 3, 12}
	if Maximo(arr) != 12 {
		t.Errorf("Maximo([10,5,8,3,12]) debe ser 12, obtuve %d", Maximo(arr))
	}
}

func TestMaximoNegativos(t *testing.T) {
	arr := []int{-5, -2, -10, -1}
	if Maximo(arr) != -1 {
		t.Errorf("Maximo([-5,-2,-10,-1]) debe ser -1, obtuve %d", Maximo(arr))
	}
}

func TestMaximoFlotantes(t *testing.T) {
	arr := []float64{1.75, 1.80, 1.65, 1.90}
	if Maximo(arr) != 1.90 {
		t.Errorf("Maximo([1.75,1.80,1.65,1.90]) debe ser 1.90, obtuve %f", Maximo(arr))
	}
}

func TestMaximoStrings(t *testing.T) {
	arr := []string{"Ana", "Luis", "Pepe", "Beatriz"}
	if Maximo(arr) != "Pepe" {
		t.Errorf("Maximo([Ana,Luis,Pepe,Beatriz]) debe ser Pepe, obtuve %s", Maximo(arr))
	}
}

func TestMinimoEnteros(t *testing.T) {
	arr := []int{10, 5, 8, 3, 12}
	if Minimo(arr) != 3 {
		t.Errorf("Minimo([10,5,8,3,12]) debe ser 3, obtuve %d", Minimo(arr))
	}
}

func TestMinimoFlotantes(t *testing.T) {
	arr := []float64{1.75, 1.80, 1.65, 1.90}
	if Minimo(arr) != 1.65 {
		t.Errorf("Minimo([1.75,1.80,1.65,1.90]) debe ser 1.65, obtuve %f", Minimo(arr))
	}
}

func TestMinimoStrings(t *testing.T) {
	arr := []string{"Ana", "Luis", "Pepe", "Beatriz"}
	if Minimo(arr) != "Ana" {
		t.Errorf("Minimo([Ana,Luis,Pepe,Beatriz]) debe ser Ana, obtuve %s", Minimo(arr))
	}
}

func TestMaximoUnElemento(t *testing.T) {
	arr := []int{42}
	if Maximo(arr) != 42 {
		t.Errorf("Maximo([42]) debe ser 42, obtuve %d", Maximo(arr))
	}
}

func TestMinimoUnElemento(t *testing.T) {
	arr := []float64{3.14}
	if Minimo(arr) != 3.14 {
		t.Errorf("Minimo([3.14]) debe ser 3.14, obtuve %f", Minimo(arr))
	}
}
