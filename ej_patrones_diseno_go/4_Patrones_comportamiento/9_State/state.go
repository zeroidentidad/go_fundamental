package state

// Interface Estado
type Estado interface {
	Escribir(string) string
}

// Estado Concreto
type EstadoNegrita struct{}

func (en *EstadoNegrita) Escribir(texto string) string {
	return "*" + texto + "*"
}

// Estado Concreto
type EstadoCursiva struct{}

func (ec *EstadoCursiva) Escribir(texto string) string {
	return "_" + texto + "_"
}

// Contexto
type EditorMarkdown struct {
	estado Estado
}

func (em *EditorMarkdown) SetEstado(estado Estado) {
	em.estado = estado
}

func (em *EditorMarkdown) Redactar(texto string) string {
	if em.estado == nil {
		return texto
	}

	return em.estado.Escribir(texto)
}
