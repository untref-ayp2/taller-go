package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("=== Crear un módulo ===")

	dir, _ := os.MkdirTemp("", "ejemplo-modulo")
	defer os.RemoveAll(dir)

	cmd := exec.Command("go", "mod", "init", "github.com/usuario/mi-modulo")
	cmd.Dir = dir
	out, _ := cmd.Output()

	fmt.Println(string(out))

	datos, _ := os.ReadFile(dir + "/go.mod")
	fmt.Println("Contenido de go.mod:")
	fmt.Println(string(datos))

	fmt.Println("El módulo se creó exitosamente.")
}
