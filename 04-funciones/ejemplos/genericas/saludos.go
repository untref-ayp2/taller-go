package genericas

import "fmt"

func Saludar(persona string) {
	fmt.Println("Hola", persona)
}

func Saludar2(persona string) string {
	saludo := "Hola " + persona
	return saludo
}
