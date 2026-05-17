package main

import "testing"

func TestLunes(t *testing.T) {
	r := diaSemana(1)
	if r != "lunes" {
		t.Errorf("diaSemana(1) = %q; esperado \"lunes\"", r)
	}
}

func TestMiercoles(t *testing.T) {
	r := diaSemana(3)
	if r != "miércoles" {
		t.Errorf("diaSemana(3) = %q; esperado \"miércoles\"", r)
	}
}

func TestDomingo(t *testing.T) {
	r := diaSemana(7)
	if r != "domingo" {
		t.Errorf("diaSemana(7) = %q; esperado \"domingo\"", r)
	}
}
