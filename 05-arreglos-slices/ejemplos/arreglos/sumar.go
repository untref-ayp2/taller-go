package arreglos

// Sumar retorna la suma de todos los elementos de un arreglo de enteros que
// recibe por parámetro
func Sumar(arreglo []int) (total int) {
	total = 0
	for i := 0; i < len(arreglo); i++ {
		total += arreglo[i]
	}
	return
}
