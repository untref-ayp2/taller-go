package main

import (
	"fmt"
	"os"
)

func main() {
	err := AgregarLinea("bitacora.txt", "nueva linea")
	if err != nil {
		fmt.Println(err)
	}
}

func AgregarLinea(archivo, linea string) error {
	f, err := os.OpenFile(archivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(linea + "\n")
	return err
}
