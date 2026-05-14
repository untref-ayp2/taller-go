package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Open("ejemplo.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	linea := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", linea, scanner.Text())
		linea++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer:", err)
	}
}
