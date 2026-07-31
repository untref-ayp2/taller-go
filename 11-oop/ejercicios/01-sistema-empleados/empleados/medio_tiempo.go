package empleados

type MedioTiempo struct {
	nombre          string
	puesto          string
	sueldoPorHora   float64
	horasTrabajadas int
}

func NewMedioTiempo(nombre string, puesto string, sueldoPorHora float64, horasTrabajadas int) *MedioTiempo {
	return nil // TODO
}

func (e *MedioTiempo) CalcularSueldo() float64 {
	return 0 // TODO
}

func (e *MedioTiempo) Puesto() string {
	return "" // TODO
}

func (e *MedioTiempo) Nombre() string {
	return "" // TODO
}
