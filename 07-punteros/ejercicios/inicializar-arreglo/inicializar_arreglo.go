package main

import "fmt"

func main() {
	var arr [5]int
	InicializarArreglo(&arr)
	fmt.Println(arr)
}

func InicializarArreglo(arr *[5]int) {
	for i := range arr {
		arr[i] = i * 2
	}
}
