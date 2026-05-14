package arreglos

// InsertionSort ordena el arreglo recibido como parametro con el métodos de
// inserción. El ordenamiento se realiza sobre el mismo arreglo
func InsertionSort(arreglo []int) {
	i := 1
	for i < len(arreglo) {
		j := i
		for j > 0 && arreglo[j-1] > arreglo[j] {
			swap(j-1, j, arreglo)
			j--
		}
		i++
	}
}

// swap intercambia los elementos en las posiciones `i` y `j` sobre el arreglo
// dado. Es decir que las modificaciones se realicen sobre el arreglo original y
// no sobre una copia
func swap(i int, j int, arreglo []int) {
	arreglo[i], arreglo[j] = arreglo[j], arreglo[i]
}
