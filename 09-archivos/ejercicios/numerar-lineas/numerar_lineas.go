package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := NumerarLineas("entrada.txt", "salida.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func NumerarLineas(entrada, salida string) error {
	f, err := os.Open(entrada)
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := os.Create(salida)
	if err != nil {
		return err
	}
	defer out.Close()

	scanner := bufio.NewScanner(f)
	linea := 1
	for scanner.Scan() {
		fmt.Fprintf(out, "%d: %s\n", linea, scanner.Text())
		linea++
	}
	return scanner.Err()
}
