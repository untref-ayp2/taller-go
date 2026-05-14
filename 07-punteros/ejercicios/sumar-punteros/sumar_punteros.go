package main

import "fmt"

func main() {
	x, y := 5, 7
	resultado := SumarPunteros(&x, &y)
	fmt.Println(resultado)
}

func SumarPunteros(a, b *int) int {
	return *a + *b
}
