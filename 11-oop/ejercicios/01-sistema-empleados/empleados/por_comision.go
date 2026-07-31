package empleados

type PorComision struct {
	nombre              string
	puesto              string
	sueldoBase          float64
	ventasRealizadas    float64
	porcentajeComision  float64
}

func NewPorComision(nombre string, puesto string, sueldoBase, ventasRealizadas, porcentajeComision float64) *PorComision {
	return nil // TODO
}

func (e *PorComision) CalcularSueldo() float64 {
	return 0 // TODO
}

func (e *PorComision) Puesto() string {
	return "" // TODO
}

func (e *PorComision) Nombre() string {
	return "" // TODO
}
