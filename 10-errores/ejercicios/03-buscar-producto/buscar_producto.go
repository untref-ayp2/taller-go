package main

import "errors"

type Producto struct {
	Codigo int
	Nombre string
}

var ErrProductoNoEncontrado = errors.New("producto no encontrado")

func BuscarProducto(productos []Producto, codigo int) (Producto, error) {
	// TODO: implementar
	return Producto{}, nil
}

func main() {
	// Usá este espacio para probar tu implementación
}
