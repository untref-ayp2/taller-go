package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = Eliminar(s, 2)
	fmt.Println(s)
}

func Eliminar(s []int, pos int) []int {
	if pos < 0 || pos >= len(s) {
		return s
	}
	return append(s[:pos], s[pos+1:]...)
}
