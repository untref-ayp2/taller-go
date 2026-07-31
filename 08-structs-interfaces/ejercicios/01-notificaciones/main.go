package main

type Email struct {
	Direccion, Mensaje string
}

func NewEmail(direccion, mensaje string) Email {
	// TODO: implementar
	return Email{}
}

func (e Email) Enviar() string {
	// TODO: implementar
	return ""
}

type SMS struct {
	Numero, Mensaje string
}

func NewSMS(numero, mensaje string) SMS {
	// TODO: implementar
	return SMS{}
}

func (s SMS) Enviar() string {
	// TODO: implementar
	return ""
}

type Notificable interface {
	Enviar() string
}

func EnviarNotificaciones(ns []Notificable) {
	// TODO: implementar
}

func main() {
	// Usá este espacio para probar tu implementación
}
