package main

import "fmt"

// --- Structs ---

type Direccion struct {
	calle, ciudad, provincia string
	numero                   uint
}

type Persona struct {
	nombre, apellido string
	edad             uint
	direccion        Direccion
}

func (p Persona) NombreCompleto() string {
	return p.nombre + " " + p.apellido
}

func (p *Persona) CumplirAnios() {
	p.edad++
}

// --- Interfaces ---

type Caminante interface {
	Avanzar(pasos int)
	Girar(grados float32)
}

func (p *Persona) Avanzar(pasos int) {
	fmt.Printf("%s avanzó %d pasos\n", p.nombre, pasos)
}

func (p *Persona) Girar(grados float32) {
	fmt.Printf("%s giró %.1f grados\n", p.nombre, grados)
}

type Perro struct {
	nombre string
}

func (perro *Perro) Avanzar(pasos int) {
	fmt.Printf("%s avanzó %d pasos\n", perro.nombre, pasos)
}

func (perro *Perro) Girar(grados float32) {
	fmt.Printf("%s giró %.1f grados\n", perro.nombre, grados)
}

func HacerCaminar(c Caminante) {
	c.Avanzar(10)
	c.Girar(90)
}

// --- Varias interfaces ---

type Trabajador interface {
	Trabajar(horas int) string
}

func (p *Persona) Trabajar(horas int) string {
	return fmt.Sprintf("%s trabajó %d horas", p.nombre, horas)
}

// --- Constructores ---

func NewPersona(nombre, apellido string, edad uint) Persona {
	return Persona{
		nombre:   nombre,
		apellido: apellido,
		edad:     edad,
	}
}

func main() {
	// Structs y acceso a campos
	p1 := Persona{nombre: "Marcelo", edad: 27}
	p2 := Persona{nombre: "Laura", apellido: "Medina", edad: 25}

	fmt.Println(p1.nombre, p1.edad)
	fmt.Println(p2.nombre, p2.apellido)

	// Métodos
	p := Persona{nombre: "Ana", apellido: "López", edad: 30}
	fmt.Println(p.NombreCompleto())
	p.CumplirAnios()
	fmt.Println(p.edad)

	// Punteros a struct
	pp := &Persona{nombre: "Laura", apellido: "Medina", edad: 25}
	fmt.Println(pp.nombre)
	pp.nombre = "María"
	fmt.Println(pp.nombre)

	// Azúcar sintáctico (receptor puntero sobre variable no puntero)
	p3 := Persona{nombre: "Ana", edad: 30}
	p3.CumplirAnios()
	fmt.Println(p3.edad)

	// Constructor New
	np := NewPersona("Laura", "Medina", 25)
	fmt.Println(np.nombre, np.apellido, np.edad)

	// Interfaces: polimorfismo
	HacerCaminar(&Persona{nombre: "Ana"})
	HacerCaminar(&Perro{nombre: "Rex"})

	// Varias interfaces
	p4 := Persona{nombre: "Laura"}
	fmt.Println(p4.Trabajar(8))

	// Variables de tipo interfaz
	ana := Persona{nombre: "Ana", apellido: "López", edad: 30}
	var c Caminante = &ana
	c.Avanzar(5)
	var t Trabajador = &ana
	fmt.Println(t.Trabajar(8))
}
