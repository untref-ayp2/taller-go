package main

import (
	"fmt"

	"github.com/untref-ayp2/taller-go/05-arreglos-slices/ejemplos/arreglos"
)

func main() {
	a1 := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3}
	a2 := a1
	a3 := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3, 4}

	fmt.Println(a1)
	arreglos.InsertionSort(a1)
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println("la suma de los elementos del arreglo es: ", arreglos.Sumar(a2))

	fmt.Println("Arreglo: ", a3)
	suma_maxima, inicial, final := arreglos.SubsecuenciaSumaMaxima(a3)
	fmt.Printf("La subsecuencia de suma maxima empieza en %v, termina en %v "+
		"y su suma es %v\n", inicial, final, suma_maxima)

	// Calcular el promedio de los elementos usando len(x)
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	fmt.Println("x: ", x)
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	// Esta instruccion dará error pues los tipos del cociente no son los mismos
	// fmt.Println(total / len(x)) (total es float64 mientras que len(x) es int)
	// len(x) debe ser convertida a float64
	fmt.Printf("Promedio de los elementos de x: %5.2f\n", total/float64(len(x)))

	// Mismo ejemplo usando otra forma de recorrer un arreglo: range
	// range itera sobre el array y value tiene el mismo significado que v[i]
	// el guión bajo reemplaza al indice del array que está siendo usado.
	total = 0
	for _, value := range x {
		total += value
	}
	fmt.Printf("Promedio range de los elementos de x: %5.2f\n", total/float64(len(x)))

	// Tajadas o Slices
	tajada := a3[inicial:]
	fmt.Println(tajada)

	for i := 0; i < len(tajada); i++ {
		tajada[i] = i
	}
	fmt.Println(tajada)
	fmt.Println(a3)

	t1 := a3[5:]
	t2 := a3[:4]
	t3 := a3[:]

	fmt.Println("t1 = ", t1)
	fmt.Println("t2 = ", t2)
	fmt.Println("t3 = ", t3)

	t3[0] = 100

	fmt.Println("a3 = ", a3)
	// append toma una tajada y le agrega elementos, devolviendo un nuevo arreglo
	t1 = append(t1, 33, 44)
	fmt.Println("t1 = ", t1)

	// con la sintaxis a3... puedo pasarle todos los elementos de un arreglo
	// a una función variádica
	t1 = append(t1, a3...)
	fmt.Println("t1 = ", t1)

	// eliminar el elemento en la posicion pos de un Slices
	fmt.Println("Eliminar el elemento en la 4ta pos")
	todos := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Todos: ", todos) //[0 1 2 3 4 5 6 7 8 9]
	remove(todos, 4)
	// todos = remove(todos, 4)
	fmt.Println("Sin el 4: ", todos) // Analizar que pasa con las dos lineas de arriba
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
