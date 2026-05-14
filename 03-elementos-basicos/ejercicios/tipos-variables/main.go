package main

import "fmt"

func main() {
	var temperatura float64 = 36.5
	const congelacion = 0

	fmt.Printf("Temperatura: %.1f°C\n", temperatura)
	fmt.Printf("Punto de congelación: %d°C\n", congelacion)
}
