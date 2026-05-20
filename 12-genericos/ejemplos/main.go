package main

import (
	"fmt"
)

// Ordenable agrupa tipos que soportan los operadores <, >, <=, >=
type Ordenable interface {
	~int | ~float64 | ~string
}

// Contiene verifica si un elemento está en el slice.
func Contiene[T comparable](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

// BuscarLineal devuelve el índice de un elemento en el slice, o -1 si no está.
func BuscarLineal[T comparable](arr []T, elem T) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}
	return -1
}

// OrdenarSeleccion ordena un slice usando el algoritmo de selección.
// La función menor determina si a debe ir antes que b.
func OrdenarSeleccion[T any](arr []T, menor func(T, T) bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if menor(arr[j], arr[minIdx]) {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Maximo devuelve el valor máximo de un slice.
func Maximo[T Ordenable](arr []T) T {
	max := arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Minimo devuelve el valor mínimo de un slice.
func Minimo[T Ordenable](arr []T) T {
	min := arr[0]
	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// MergeSlices combina dos slices ordenados en uno solo ordenado.
// La función menor determina si a debe ir antes que b.
func MergeSlices[T any](a, b []T, menor func(T, T) bool) []T {
	result := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if menor(a[i], b[j]) {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func main() {
	// --- Contiene ---
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("Contiene 30:", Contiene(numeros, 30))   // true
	fmt.Println("Contiene 99:", Contiene(numeros, 99))   // false

	nombres := []string{"Ana", "Luis", "Pepe"}
	fmt.Println("Contiene Luis:", Contiene(nombres, "Luis")) // true

	// --- BuscarLineal ---
	fmt.Println("Posicion de 30:", BuscarLineal(numeros, 30)) // 2
	fmt.Println("Posicion de 99:", BuscarLineal(numeros, 99)) // -1

	// --- OrdenarSeleccion ---
	floats := []float64{3.5, 1.2, 4.8, 2.1}
	OrdenarSeleccion(floats, func(a, b float64) bool { return a < b })
	fmt.Println("Ordenado:", floats) // [1.2 2.1 3.5 4.8]

	// Con structs
	type Persona struct {
		Nombre string
		Edad   int
	}
	personas := []Persona{
		{"Ana", 30},
		{"Luis", 25},
		{"Pepe", 35},
	}
	OrdenarSeleccion(personas, func(a, b Persona) bool { return a.Edad < b.Edad })
	fmt.Println("Personas por edad:", personas) // [{Luis 25} {Ana 30} {Pepe 35}]

	// --- Maximo ---
	enteros := []int{10, 5, 8, 3, 12}
	fmt.Println("Maximo:", Maximo(enteros)) // 12

	alturas := []float64{1.75, 1.80, 1.65, 1.90}
	fmt.Println("Maximo:", Maximo(alturas)) // 1.9

	palabras := []string{"Ana", "Luis", "Pepe", "Beatriz"}
	fmt.Println("Maximo:", Maximo(palabras)) // Pepe (lexicográfico)

	// --- Minimo ---
	fmt.Println("Minimo:", Minimo(enteros))   // 3
	fmt.Println("Minimo:", Minimo(alturas))   // 1.65

	// --- MergeSlices ---
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	mezclado := MergeSlices(a, b, func(x, y int) bool { return x < y })
	fmt.Println("Mezclado:", mezclado) // [1 2 3 4 5 6 7 8]
}
