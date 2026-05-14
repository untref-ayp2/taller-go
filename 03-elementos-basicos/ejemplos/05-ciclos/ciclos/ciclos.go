package ciclos

import "fmt"

// Java itera utilizando el clásico `for` que ya conocemos de Java
func Java() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t", i)
	}
	fmt.Println()
}

// While itera utilizando el ciclo `for` como si fuera un `while`
func While() {
	i := 0
	for i < 10 {
		fmt.Printf("%v\t", i)
		i++
	}
	fmt.Println()
}

// Infinito itera sin especificar una condición, el fin del ciclo se realiza con el uso de `break`
func Infinito() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Printf("%v\t", i)
		i++
	}
	fmt.Println()
}

// RangeArreglo itera sobre un arreglo utilizando el constructor `range`
func RangeArreglo() {
	arreglo := [10]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	for indice, valor := range arreglo {
		fmt.Printf("2**%v = %v\n", indice, valor)
	}
	fmt.Println()
}

// RangeString itera sobre un string utilizando el constructor `range`
func RangeString() {
	cadena := "Algoritmos y Programación II"
	for _, letra := range cadena {
		fmt.Printf("%c ", letra)
	}
	fmt.Println()
}
