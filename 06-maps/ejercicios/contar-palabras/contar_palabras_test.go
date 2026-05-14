package main

import "testing"

func TestContarPalabras(t *testing.T) {
	resultado := ContarPalabras("hola mundo hola")
	if resultado["hola"] != 2 {
		t.Errorf("hola aparece 2 veces, obtuve %d", resultado["hola"])
	}
	if resultado["mundo"] != 1 {
		t.Errorf("mundo aparece 1 vez, obtuve %d", resultado["mundo"])
	}
}

func TestContarPalabrasVacio(t *testing.T) {
	resultado := ContarPalabras("")
	if len(resultado) != 0 {
		t.Errorf("texto vacio deberia devolver mapa vacio, obtuve %v", resultado)
	}
}
