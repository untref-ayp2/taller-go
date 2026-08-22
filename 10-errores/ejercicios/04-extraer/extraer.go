package main

type SaldoInsuficienteError struct {
	Saldo float64
	Monto float64
}

// Error implementa la interfaz error.
func (e SaldoInsuficienteError) Error() string {
	return "" // TODO: devolver un mensaje descriptivo usando Saldo y Monto
}

func Extraer(saldo, monto float64) (float64, error) {
	// TODO: implementar
	return 0, nil
}

func main() {
	// Usá este espacio para probar tu implementación
}
