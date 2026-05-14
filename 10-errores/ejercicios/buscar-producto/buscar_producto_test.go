package main

import "testing"

func TestBuscarProductoOK(t *testing.T) {
	p := []Producto{{1, "A"}, {2, "B"}}
	r, err := BuscarProducto(p, 1)
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	if r.Nombre != "A" {
		t.Errorf("producto 1 -> %q; esperado A", r.Nombre)
	}
}

func TestBuscarProductoNoEncontrado(t *testing.T) {
	p := []Producto{{1, "A"}}
	_, err := BuscarProducto(p, 99)
	if err == nil {
		t.Error("producto inexistente deberia devolver error")
	}
}
