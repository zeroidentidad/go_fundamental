package interpreter

// Contexto
type Contexto struct {
	Palabra string
}

// Interface
type ExpresionAbstracta interface {
	Interpretar(Contexto) bool
}

// Expresion Terminal
type ExpresionTerminal struct {
	Palabra string
}

func (et *ExpresionTerminal) Interpretar(contexto Contexto) bool {
	return contexto.Palabra == et.Palabra
}

// Expresion No Terminal
type ExpresionOR struct {
	expresionA ExpresionAbstracta
	expresionB ExpresionAbstracta
}

func (eo *ExpresionOR) Interpretar(contexto Contexto) bool {
	return eo.expresionA.Interpretar(contexto) || eo.expresionB.Interpretar(contexto)
}

// Expresion No Terminal
type ExpresionAND struct {
	expresionA ExpresionAbstracta
	expresionB ExpresionAbstracta
}

func (ea *ExpresionAND) Interpretar(contexto Contexto) bool {
	return ea.expresionA.Interpretar(contexto) && ea.expresionB.Interpretar(contexto)
}
