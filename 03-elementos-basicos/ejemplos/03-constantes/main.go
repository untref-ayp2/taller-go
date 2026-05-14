package main

import (
	"fmt"
	"math"
)

const s string = "constante"

func main() {
	// Constantes predefinidas
	fmt.Println(math.Pi)
	fmt.Println(math.E)

	// Constantes de usuario
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(s)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	// ERROR:
	// n = 1
	// s = "h"
}
