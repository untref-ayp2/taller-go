package main

type SaldoInsuficienteError struct {
	Saldo float64
	Monto float64
}

// TODO: implementar el método Error() para que SaldoInsuficienteError implemente error

func Extraer(saldo, monto float64) (float64, error) {
	// TODO: implementar
	return 0, nil
}

func main() {
	// Usá este espacio para probar tu implementación
}
