package main

// Ordenable agrupa tipos que soportan operadores de comparación (<, >, etc.)
type Ordenable interface {
	~int | ~float64 | ~string
}

// Maximo devuelve el valor máximo de un slice.
// El slice debe tener al menos un elemento.
func Maximo[T Ordenable](arr []T) T {
	// TODO: implementar
	var zero T
	return zero
}

// Minimo devuelve el valor mínimo de un slice.
// El slice debe tener al menos un elemento.
func Minimo[T Ordenable](arr []T) T {
	// TODO: implementar
	var zero T
	return zero
}
