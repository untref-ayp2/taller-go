package main

import (
	"strings"
	"testing"
)

func TestNewEmail(t *testing.T) {
	e := NewEmail("a@b.com", "hola")
	if e.Direccion != "a@b.com" || e.Mensaje != "hola" {
		t.Errorf("NewEmail no creó correctamente el Email")
	}
}

func TestEmailEnviar(t *testing.T) {
	e := Email{Direccion: "a@b.com", Mensaje: "hola"}
	r := e.Enviar()
	if !strings.Contains(r, "a@b.com") || !strings.Contains(r, "hola") {
		t.Errorf("Enviar no incluye dirección ni mensaje: %s", r)
	}
}

func TestNewSMS(t *testing.T) {
	s := NewSMS("1234", "hola")
	if s.Numero != "1234" || s.Mensaje != "hola" {
		t.Errorf("NewSMS no creó correctamente el SMS")
	}
}

func TestSMSEnviar(t *testing.T) {
	s := SMS{Numero: "1234", Mensaje: "hola"}
	r := s.Enviar()
	if !strings.Contains(r, "1234") || !strings.Contains(r, "hola") {
		t.Errorf("Enviar no incluye número ni mensaje: %s", r)
	}
}

func TestNotificableInterface(t *testing.T) {
	var n Notificable
	n = Email{Direccion: "a@b.com", Mensaje: "test"}
	_ = n.Enviar()
	n = SMS{Numero: "1234", Mensaje: "test"}
	_ = n.Enviar()
}

func TestEnviarNotificaciones(t *testing.T) {
	ns := []Notificable{
		Email{Direccion: "a@b.com", Mensaje: "hola"},
		SMS{Numero: "1234", Mensaje: "mundo"},
	}
	EnviarNotificaciones(ns)
}
