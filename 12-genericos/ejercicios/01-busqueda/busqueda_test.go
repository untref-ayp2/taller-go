package main

import "testing"

func TestContieneEnteros(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Contiene(arr, 3) {
		t.Error("Contiene([1,2,3,4,5], 3) debe ser true")
	}
	if Contiene(arr, 10) {
		t.Error("Contiene([1,2,3,4,5], 10) debe ser false")
	}
}

func TestContieneStrings(t *testing.T) {
	arr := []string{"Ana", "Luis", "Pepe"}
	if !Contiene(arr, "Luis") {
		t.Error("Contiene([Ana, Luis, Pepe], Luis) debe ser true")
	}
	if Contiene(arr, "Maria") {
		t.Error("Contiene([Ana, Luis, Pepe], Maria) debe ser false")
	}
}

func TestContieneVacio(t *testing.T) {
	arr := []int{}
	if Contiene(arr, 1) {
		t.Error("Contiene([], 1) debe ser false")
	}
}

func TestPosicionEnteros(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}
	if Posicion(arr, 30) != 2 {
		t.Error("Posicion([10,20,30,40,50], 30) debe ser 2")
	}
	if Posicion(arr, 99) != -1 {
		t.Error("Posicion([10,20,30,40,50], 99) debe ser -1")
	}
}

func TestPosicionStrings(t *testing.T) {
	arr := []string{"Ana", "Luis", "Pepe"}
	if Posicion(arr, "Pepe") != 2 {
		t.Error("Posicion([Ana, Luis, Pepe], Pepe) debe ser 2")
	}
	if Posicion(arr, "Zoe") != -1 {
		t.Error("Posicion([Ana, Luis, Pepe], Zoe) debe ser -1")
	}
}

func TestPosicionVacio(t *testing.T) {
	arr := []string{}
	if Posicion(arr, "a") != -1 {
		t.Error("Posicion([], a) debe ser -1")
	}
}
