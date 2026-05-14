package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	archivo := "numeros.txt"
	fmt.Println(SumarNumeros(archivo))
}

func SumarNumeros(archivo string) int {
	f, err := os.Open(archivo)
	if err != nil {
		return 0
	}
	defer f.Close()
	suma := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err == nil {
			suma += n
		}
	}
	return suma
}
