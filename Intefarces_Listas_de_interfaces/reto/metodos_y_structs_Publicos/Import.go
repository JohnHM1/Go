package publicos

type Circulo struct {
	Pi    float64
	Radio float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Cuadrado struct {
	Base float64
}

func (c Cuadrado) Area1() float64 {
	return c.Base * c.Base
}

func (r Rectangulo) Area1() float64 {
	return r.Altura * r.Base
}

func (ci Circulo) Area1() float64 {
	return ci.Pi * (ci.Radio * ci.Radio)
}
