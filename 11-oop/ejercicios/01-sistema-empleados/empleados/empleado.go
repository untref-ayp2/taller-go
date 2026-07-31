package empleados

type Empleado interface {
	CalcularSueldo() float64
	Puesto() string
	Nombre() string
}
