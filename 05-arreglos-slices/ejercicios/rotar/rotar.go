package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	Rotar(s, 2)
	fmt.Println(s)
}

func Rotar(s []int, k int) {
	if len(s) == 0 {
		return
	}
	k = k % len(s)
	Invertir(s[:k])
	Invertir(s[k:])
	Invertir(s)
}

func Invertir(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
