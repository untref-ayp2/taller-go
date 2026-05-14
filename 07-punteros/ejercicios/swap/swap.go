package main

import "fmt"

func main() {
	x, y := 10, 20
	Swap(&x, &y)
	fmt.Println(x, y)
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}
