package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
	"time"
)

// Declaraciones "globales"
var entero = 1
var enteroDeclarado int = 1
var entero1, entero2, entero3 = 1, 2, 3
var entero4, entero5, entero6 int = 1, 2, 3
var enteroZ1, enteroZ2, enteroZ3 int

// No es posible usar la declaración corta fuera de funciones
// otroEntero := 1

// var nombre string, apellido string, edad int = "Nombre", "Apellido", 120
var nombre, apellido, edad = "Nombre", "Apellido", 120

// Declaración múltiple.
var (
	encendido bool       = false
	MaxUInt64 uint64     = 1<<64 - 1
	complejoZ complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//enteroZ1 = 10
	//enteroAsignacionCortaEImplicita := 1
	//var enteroAsignacionImplicita = 1
	//var enteroAsignacionExplicita int = 1
	//var enteroVacioExplicito int // = 0 -> Valor cero en su tipo
	//var enteroAsignacionExplicita int = 1.2 // ERROR

	//var f float32 = 1.2
	//var enteroAsignacionExplicita int = int(f)

	//entero1, entero2 := 1, 2

	// Variable local "esconde" a la variable "global"
	//enteroZ1 := 1 // OK

	// No se puede definir nuevamente, como toda variable (en su contexto)
	//enteroZ1 := 2 // ERROR

	fmt.Println(enteroZ1)
	fmt.Println("Contenido de los enteros sin asignar:", enteroZ1, enteroZ2, enteroZ3)

	fmt.Printf("Tipo: %-10T Valor: %v\n", encendido, encendido)
	fmt.Printf("Tipo: %-10T Valor: %v\n", MaxUInt64, MaxUInt64)
	fmt.Printf("Tipo: %-10T Valor: %v\n", complejoZ, complejoZ)

	currentTime := time.Now()
	fmt.Println("El tipo de `time.Now()` es:", reflect.TypeOf(currentTime))
	fmt.Println("El tipo de `complejoZ` es:", reflect.TypeOf(complejoZ))
}
