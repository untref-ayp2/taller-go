/*
Tipos de datos:

	bool       // true o false
	string     // "Hola mundo"
	int        // 32 o 64 bits
	int8       // -128 a 127
	int16      // -32768 a 32767
	int32      // -2147483648 a 2147483647
	int64      // -9223372036854775808 a 9223372036854775807
	uint       // 32 o 64 bits
	uint8      // 0 a 255
	uint16     // 0 a 65535
	uint32     // 0 a 4294967295
	uint64     // 0 a 18446744073709551615
	float32    // 32 bits
	float64    // 64 bits
	complex64  // 32 bits
	complex128 // 64 bits
	byte       // alias para uint8
	rune       // alias para int32, representa una posición en código Unicode
	uintptr    // utilizado para guardar una dirección de puntero

Nota: Los tamaños de int, uint y uintptr varian segun el sistema (normalmente 32
bits en SO de 32 bits y 64 bits en caso de SO de 64 bits)
*/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var i int
	fmt.Printf("Tamaño en bytes del tipo int: %d\n", reflect.TypeOf(i).Size())
	fmt.Printf("Tamaño en bytes de la variable i: %d\n", unsafe.Sizeof(i))
}
