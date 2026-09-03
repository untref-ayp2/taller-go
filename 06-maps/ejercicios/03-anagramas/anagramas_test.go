package main

import "testing"

func TestSonAnagramasSi(t *testing.T) {
	if !SonAnagramas("listen", "silent") {
		t.Errorf("listen y silent son anagramas")
	}
}

func TestSonAnagramasNo(t *testing.T) {
	if SonAnagramas("hola", "mundo") {
		t.Errorf("hola y mundo no son anagramas")
	}
}

func TestSonAnagramasCaseInsensitive(t *testing.T) {
	if !SonAnagramas("Listen", "Silent") {
		t.Errorf("Listen y Silent son anagramas (case insensitive)")
	}
}

func TestSonAnagramasDistintoLargo(t *testing.T) {
	if SonAnagramas("abc", "abcd") {
		t.Errorf("distinto largo no puede ser anagrama")
	}
}

func TestNoSonAnagramasMismoLargo(t *testing.T) {
	if SonAnagramas("abc", "abd") {
		t.Errorf("abc y abd no son anagramas a pesar de tener la misma longitud")
	}
}

func TestVacias(t *testing.T) {
	if !SonAnagramas("", "") {
		t.Errorf("dos strings vacios son anagramas")
	}
}

func TestUnaVaciaUnaNo(t *testing.T) {
	if SonAnagramas("abc", "") {
		t.Errorf("abc y string vacio no pueden ser anagramas")
	}
}
