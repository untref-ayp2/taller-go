# Ejercicio: Sistema de empleados

Implementá un programa que modele distintos tipos de empleados de una empresa.

## Interfaz `Empleado`

La interfaz `Empleado` ya está definida en el paquete `empleados/`:

- `CalcularSueldo() float64`
- `Puesto() string`
- `Nombre() string`

## Qué implementar

### a. TiempoCompleto

Campos: `nombre`, `puesto`, `sueldoMensual`.
Constructor: `NewTiempoCompleto(nombre, puesto string, sueldoMensual float64) *TiempoCompleto`
Sueldo: `sueldoMensual`.

### b. MedioTiempo

Campos: `nombre`, `puesto`, `sueldoPorHora`, `horasTrabajadas`.
Constructor: `NewMedioTiempo(nombre, puesto string, sueldoPorHora float64, horasTrabajadas int) *MedioTiempo`
Sueldo: `sueldoPorHora * horasTrabajadas`.

### c. PorComision

Campos: `nombre`, `puesto`, `sueldoBase`, `ventasRealizadas`, `porcentajeComision`.
Constructor: `NewPorComision(nombre, puesto string, sueldoBase, ventasRealizadas, porcentajeComision float64) *PorComision`
Sueldo: `sueldoBase + ventasRealizadas * porcentajeComision`.

### d. Gerente

**Embebe** `TiempoCompleto` y agrega un campo `bono`.
Constructor: `NewGerente(nombre string, sueldoMensual float64, bono float64) *Gerente`
Sueldo: `sueldoMensual + bono`.
Puesto: siempre devuelve `"Gerente"`.
El resto de los métodos (`Nombre`, `CalcularSueldo` de `TiempoCompleto`) se heredan por embedding — salvo `CalcularSueldo` que se sobreescribe para incluir el bono.

### e. Funciones

En `main.go`:

- `TotalSueldos(empleados []Empleado) float64` — suma los sueldos de todos los empleados.
- `EmpleadoConMasResponsabilidad(empleados []Empleado) string` — devuelve el nombre del empleado con mayor sueldo.

### f. Función main

Creá al menos un empleado de cada tipo, calculá el total y mostrá quién tiene el mayor sueldo.
