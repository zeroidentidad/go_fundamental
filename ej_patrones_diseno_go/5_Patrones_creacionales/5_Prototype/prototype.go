package prototype

// Prototipo Interface
type Prototipo interface {
	Clonar() *Elemento
}

// Prototipo Concreto
type Elemento struct {
	Material string
	Copias   int
}

func (e *Elemento) Clonar() *Elemento {
	return &Elemento{
		Material: e.Material,
		Copias:   e.Copias + 1,
	}
}
