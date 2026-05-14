package main

import "fmt"

func main() {
	fmt.Println(Promedio([]float64{4.5, 7.0, 8.5, 6.0}))
}

func Promedio(notas []float64) float64 {
	if len(notas) == 0 {
		return 0
	}
	suma := 0.0
	for _, n := range notas {
		suma += n
	}
	return suma / float64(len(notas))
}
