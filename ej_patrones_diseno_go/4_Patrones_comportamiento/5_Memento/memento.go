package memento

// Interface
type Memento interface {
	SetContenido(string)
	GetContenido() string
}

// Memento
type EditorMemento struct {
	contenido string
}

func (em *EditorMemento) SetContenido(contenido string) {
	em.contenido = contenido
}

func (em *EditorMemento) GetContenido() string {
	return em.contenido
}

// Originador
type Editor struct {
	contenido string
}

func (e *Editor) VerContenido() string {
	return e.contenido
}

func (e *Editor) Escribir(texto string) {
	e.contenido = e.contenido + " " + texto
}

func (e *Editor) Guardar() Memento {
	editorMemento := &EditorMemento{}
	editorMemento.SetContenido(e.contenido)

	return editorMemento
}

func (e *Editor) Restaurar(memento Memento) {
	e.contenido = memento.GetContenido()
}
