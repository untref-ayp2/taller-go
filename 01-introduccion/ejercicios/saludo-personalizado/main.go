package main

import "fmt"

func main() {
	nombre := "mundo"
	fmt.Println(saludar(nombre))
}

func saludar(nombre string) string {
	return "¡Hola, " + nombre + "!"
}
