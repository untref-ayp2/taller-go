package main

import "testing"

func TestConcatenar(t *testing.T) {
	s := "Hola"
	Concatenar(&s, " Mundo")
	if s != "Hola Mundo" {
		t.Errorf("Concatenar(Hola, Mundo) = %q; esperado %q", s, "Hola Mundo")
	}
}

func TestConcatenarVacio(t *testing.T) {
	s := ""
	Concatenar(&s, "Go")
	if s != "Go" {
		t.Errorf("Concatenar(vacio, Go) = %q; esperado %q", s, "Go")
	}
}

func TestConcatenarSufijoVacio(t *testing.T) {
	s := "Hola"
	Concatenar(&s, "")
	if s != "Hola" {
		t.Errorf("Concatenar(Hola, vacio) = %q; esperado %q", s, "Hola")
	}
}
