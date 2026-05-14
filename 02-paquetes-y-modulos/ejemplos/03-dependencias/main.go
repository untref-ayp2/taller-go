package main

import (
	"fmt"
)

func main() {
	fmt.Println("Para agregar una dependencia externa:")
	fmt.Println("  go get github.com/google/uuid")
	fmt.Println("")
	fmt.Println("Para sincronizar el archivo go.mod:")
	fmt.Println("  go mod tidy")
	fmt.Println("")
	fmt.Println("Para ver las dependencias del módulo:")
	fmt.Println("  go list -m all")
}
