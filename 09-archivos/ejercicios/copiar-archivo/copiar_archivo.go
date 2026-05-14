package main

import (
	"fmt"
	"os"
)

func main() {
	err := CopiarArchivo("origen.txt", "destino.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func CopiarArchivo(origen, destino string) error {
	datos, err := os.ReadFile(origen)
	if err != nil {
		return err
	}
	return os.WriteFile(destino, datos, 0644)
}
