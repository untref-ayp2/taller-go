package empleados

type Gerente struct {
	TiempoCompleto
	bono float64
}

func NewGerente(nombre string, sueldoMensual float64, bono float64) *Gerente {
	return nil // TODO
}

func (e *Gerente) CalcularSueldo() float64 {
	return 0 // TODO
}

func (e *Gerente) Puesto() string {
	return "Gerente"
}
