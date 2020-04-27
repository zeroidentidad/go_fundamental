package visitor

// Visitante
type Visitante interface {
	VisitarSuperpoder(*ElementoSuperpoder)
	VisitarArma(*ElementoArma)
}

// Visitante Concreto
type VisitanteNivel1 struct{}

func (v *VisitanteNivel1) VisitarSuperpoder(elementoSuperpoder *ElementoSuperpoder) {
	elementoSuperpoder.Poder = "Rayo superpoderoso"
}

func (v *VisitanteNivel1) VisitarArma(elementoArma *ElementoArma) {
	elementoArma.Arma = "Espada de dos manos"
}

// Visitante Concreto
type VisitanteNivel0 struct{}

func (v *VisitanteNivel0) VisitarSuperpoder(elementoSuperpoder *ElementoSuperpoder) {
	elementoSuperpoder.Poder = "Rayo simple"
}

func (v *VisitanteNivel0) VisitarArma(elementoArma *ElementoArma) {
	elementoArma.Arma = "Espada de una mano"
}

// Elemento
type Elemento interface {
	Aceptar(Visitante)
}

// Elemento Concreto
type ElementoSuperpoder struct {
	Poder string
}

func (e *ElementoSuperpoder) Aceptar(visitante Visitante) {
	visitante.VisitarSuperpoder(e)
}

// Elemento Concreto
type ElementoArma struct {
	Arma string
}

func (e *ElementoArma) Aceptar(visitante Visitante) {
	visitante.VisitarArma(e)
}
