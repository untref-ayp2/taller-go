package main

import (
	"os"
	"testing"
)

func TestSumarNumeros(t *testing.T) {
	contenido := []byte("1\n2\n3\n4\n5\n")
	os.WriteFile("_nums.txt", contenido, 0644)
	defer os.Remove("_nums.txt")

	r := SumarNumeros("_nums.txt")
	if r != 15 {
		t.Errorf("SumarNumeros = %d; esperado 15", r)
	}
}
