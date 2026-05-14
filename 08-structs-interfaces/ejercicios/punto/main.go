package main

import "fmt"

func main() {
	p := NuevoPunto(3, 4)
	fmt.Println(p)
	fmt.Println("Distancia al origen:", p.DistanciaOrigen())
}

type Punto struct {
	X, Y float64
}

func NuevoPunto(x, y float64) Punto {
	return Punto{X: x, Y: y}
}

func (p Punto) DistanciaOrigen() float64 {
	return p.X*p.X + p.Y*p.Y
}
