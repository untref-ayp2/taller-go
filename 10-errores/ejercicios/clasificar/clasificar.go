package main

import (
	"errors"
	"fmt"
)

func main() {
	c, err := ClasificarEdad(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}

func ClasificarEdad(edad int) (string, error) {
	if edad < 0 {
		return "", errors.New("edad no puede ser negativa")
	}
	if edad < 13 {
		return "nino", nil
	} else if edad < 18 {
		return "adolescente", nil
	} else {
		return "adulto", nil
	}
}
