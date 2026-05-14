package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo := "texto.txt"
	fmt.Println(ContarLineas(archivo))
}

func ContarLineas(archivo string) int {
	f, err := os.Open(archivo)
	if err != nil {
		return 0
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineas := 0
	for scanner.Scan() {
		lineas++
	}
	return lineas
}
