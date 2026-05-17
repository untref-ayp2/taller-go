package main

import "testing"

func TestSumaUnoACien(t *testing.T) {
	r := sumarUnoACien()
	if r != 5050 {
		t.Errorf("sumarUnoACien() = %d; esperado 5050", r)
	}
}
