package main

import (
	"errors"
	"fmt"
)

func main() {
	s := []int{10, 20, 30, 40, 50}
	r, err := Extraer(s, 1, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Extraer(s []int, inicio, fin int) ([]int, error) {
	if inicio < 0 || fin > len(s) || inicio > fin {
		return nil, errors.New("indices invalidos")
	}
	return append([]int{}, s[inicio:fin]...), nil
}
