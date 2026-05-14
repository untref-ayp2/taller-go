package main

import (
	"errors"
	"fmt"
)

type Producto struct {
	Codigo int
	Nombre string
}

func main() {
	productos := []Producto{
		{1, "Manzana"},
		{2, "Pera"},
		{3, "Banana"},
	}
	p, err := BuscarProducto(productos, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p.Nombre)
	}
}

func BuscarProducto(productos []Producto, codigo int) (Producto, error) {
	for _, p := range productos {
		if p.Codigo == codigo {
			return p, nil
		}
	}
	return Producto{}, errors.New("producto no encontrado")
}
