package condicionales

import (
	"fmt"
	"time"
)

func Condicional(numero int) {
	if numero%2 == 0 {
		fmt.Printf("%v es par\n", numero)
	} else {
		fmt.Printf("%v es impar\n", numero)
	}

	if numero%4 == 0 {
		fmt.Printf("%v es divisible por 4\n", numero)
	}

	if numero < 0 {
		fmt.Printf("%v es negativo\n", numero)
	}

	if (numero > 0 && numero < 10) || (numero < 0 && numero > -10) {
		fmt.Printf("%v tiene un dígito\n", numero)
	} else {
		fmt.Printf("%v tiene varios dígitos\n", numero)
	}
}

func SwtichBasico() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hoy es fin de semana")
	default:
		fmt.Println("Hoy es un día laborable")
	}
}

func SwitchSinCondicion(hora int) {
	switch { // no hay expresion booleana
	case hora < 12:
		fmt.Printf("Son las %v buen día\n", hora)
	case hora < 19:
		fmt.Printf("Buenas tardes, son las %v\n", hora)
	default:
		fmt.Printf("Buenas noches, son las %v\n", hora)
	}
}

func SwitchMultiple(caracter rune) {
	switch caracter {
	case ' ', '\t', '\n', '\f', '\r':
		fmt.Println("Es blanco")
		return
	}
	fmt.Println("No es un blanco")
}

func SwitchFallthrough(numero int) {
	switch numero {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}
