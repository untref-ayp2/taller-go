package main

import (
	"testing"

	"github.com/untref-ayp2/taller-go/11-oop/ejercicios/sistema-empleados/empleados"
)

func TestTiempoCompletoSueldo(t *testing.T) {
	e := empleados.NewTiempoCompleto("Ana", "Analista", 50000)
	if e.CalcularSueldo() != 50000 {
		t.Errorf("TiempoCompleto sueldo: esperado 50000, obtenido %f", e.CalcularSueldo())
	}
}

func TestMedioTiempoSueldo(t *testing.T) {
	e := empleados.NewMedioTiempo("Luis", "Asistente", 1000, 80)
	if e.CalcularSueldo() != 80000 {
		t.Errorf("MedioTiempo sueldo: esperado 80000, obtenido %f", e.CalcularSueldo())
	}
}

func TestPorComisionSueldo(t *testing.T) {
	e := empleados.NewPorComision("Carla", "Vendedora", 20000, 100000, 0.15)
	if e.CalcularSueldo() != 35000 {
		t.Errorf("PorComision sueldo: esperado 35000, obtenido %f", e.CalcularSueldo())
	}
}

func TestGerenteSueldo(t *testing.T) {
	e := empleados.NewGerente("Pedro", 80000, 20000)
	if e.CalcularSueldo() != 100000 {
		t.Errorf("Gerente sueldo: esperado 100000, obtenido %f", e.CalcularSueldo())
	}
}

func TestGerenteEmbebeTiempoCompleto(t *testing.T) {
	e := empleados.NewGerente("Pedro", 80000, 20000)
	if e == nil {
		t.Fatal("NewGerente no debería devolver nil")
	}
	if e.Nombre() != "Pedro" {
		t.Errorf("Gerente.Nombre(): esperado Pedro, obtenido %s", e.Nombre())
	}
	if e.Puesto() != "Gerente" {
		t.Errorf("Gerente.Puesto(): esperado Gerente, obtenido %s", e.Puesto())
	}
}

func TestTotalSueldos(t *testing.T) {
	empleados := []empleados.Empleado{
		empleados.NewTiempoCompleto("Ana", "Analista", 50000),
		empleados.NewMedioTiempo("Luis", "Asistente", 1000, 80),
		empleados.NewPorComision("Carla", "Vendedora", 20000, 100000, 0.15),
		empleados.NewGerente("Pedro", 80000, 20000),
	}
	total := TotalSueldos(empleados)
	esperado := 50000.0 + 80000.0 + 35000.0 + 100000.0
	if total != esperado {
		t.Errorf("TotalSueldos: esperado %f, obtenido %f", esperado, total)
	}
}

func TestEmpleadoConMasResponsabilidad(t *testing.T) {
	empleados := []empleados.Empleado{
		empleados.NewTiempoCompleto("Ana", "Analista", 50000),
		empleados.NewMedioTiempo("Luis", "Asistente", 1000, 80),
		empleados.NewPorComision("Carla", "Vendedora", 20000, 100000, 0.15),
		empleados.NewGerente("Pedro", 80000, 20000),
	}
	nombre := EmpleadoConMasResponsabilidad(empleados)
	if nombre != "Pedro" {
		t.Errorf("EmpleadoConMasResponsabilidad: esperado Pedro, obtenido %s", nombre)
	}
}

func TestEmpleadoInterface(t *testing.T) {
	var e empleados.Empleado

	e = empleados.NewTiempoCompleto("Ana", "Analista", 50000)
	if _, ok := e.(*empleados.TiempoCompleto); !ok {
		t.Error("TiempoCompleto no implementa Empleado")
	}

	e = empleados.NewMedioTiempo("Luis", "Asistente", 1000, 80)
	if _, ok := e.(*empleados.MedioTiempo); !ok {
		t.Error("MedioTiempo no implementa Empleado")
	}

	e = empleados.NewPorComision("Carla", "Vendedora", 20000, 100000, 0.15)
	if _, ok := e.(*empleados.PorComision); !ok {
		t.Error("PorComision no implementa Empleado")
	}

	e = empleados.NewGerente("Pedro", 80000, 20000)
	if _, ok := e.(*empleados.Gerente); !ok {
		t.Error("Gerente no implementa Empleado")
	}
}
