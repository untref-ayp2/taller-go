package main

import (
	"fmt"
	"os"
)

func main() {
	datos, err := os.ReadFile("ejemplo.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	fmt.Println("Contenido del archivo:")
	fmt.Println(string(datos))
}
