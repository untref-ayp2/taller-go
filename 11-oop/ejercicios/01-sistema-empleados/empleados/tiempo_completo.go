package empleados

type TiempoCompleto struct {
	nombre       string
	puesto       string
	sueldoMensual float64
}

func NewTiempoCompleto(nombre string, puesto string, sueldoMensual float64) *TiempoCompleto {
	return nil // TODO
}

func (e *TiempoCompleto) CalcularSueldo() float64 {
	return 0 // TODO
}

func (e *TiempoCompleto) Puesto() string {
	return "" // TODO
}

func (e *TiempoCompleto) Nombre() string {
	return "" // TODO
}
