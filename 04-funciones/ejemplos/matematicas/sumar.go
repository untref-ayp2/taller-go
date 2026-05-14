package matematicas

// Sumar es una función variádica, es decir con una cantidad
// indefinida de enteros
func Sumar(numeros ...int) (total int) {
	// total = 0
	for _, valor := range numeros {
		total += valor
	}
	return
}
