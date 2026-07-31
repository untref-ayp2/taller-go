package main

import (
	"fmt"
	"os"
)

func main() {
	contenido := []byte("Hola, archivo!\nEsta es una línea nueva.\n")

	err := os.WriteFile("salida.txt", contenido, 0644)
	if err != nil {
		fmt.Println("Error al escribir:", err)
		return
	}
	fmt.Println("Archivo escrito exitosamente.")

	archivo, err := os.Create("otro.txt")
	if err != nil {
		fmt.Println("Error al crear:", err)
		return
	}
	defer archivo.Close()

	fmt.Fprintln(archivo, "Escrito con os.Create")
	fmt.Println("Archivo creado con os.Create.")
}
