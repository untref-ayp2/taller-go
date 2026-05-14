package main

import "fmt"

func main() {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"a": 1, "b": 2}
	fmt.Println(Igual(a, b))
}

func Igual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok || va != vb {
			return false
		}
	}
	return true
}
