package main

import "testing"

func TestExtraerOK(t *testing.T) {
	r, err := Extraer(100.0, 40.0)
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	if r != 60.0 {
		t.Errorf("Extraer(100,40) = %f; esperado 60", r)
	}
}

func TestExtraerSaldoInsuficiente(t *testing.T) {
	_, err := Extraer(50.0, 100.0)
	if err == nil {
		t.Fatal("deberia devolver SaldoInsuficienteError")
	}
	sie, ok := err.(SaldoInsuficienteError)
	if !ok {
		t.Errorf("error deberia ser SaldoInsuficienteError, obtuve %T", err)
	}
	if sie.Saldo != 50.0 || sie.Monto != 100.0 {
		t.Errorf("SaldoInsuficienteError.Saldo=%f .Monto=%f; esperado 50, 100", sie.Saldo, sie.Monto)
	}
}
