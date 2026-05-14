package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := Dividir(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division por cero")
	}
	return a / b, nil
}
