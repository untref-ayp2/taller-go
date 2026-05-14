package figuras

import "fmt"

type Punto struct {
	x, y int
}

func NewPunto(x, y int) Punto {
	return Punto{x: x, y: y}
}

func (p *Punto) Mover(x, y int) {
	p.x += x
	p.y += y
}

func (p Punto) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}
