package iterator

// Iterador Interface
type Iterador interface {
	Valor() string
	Siguiente()
	Anterior()
}

// Agregado Interface
type Agregado interface {
	CrearIterador() Iterador
}

// Agregado Concreto
type Radio struct {
	Estaciones []string
}

func (r *Radio) CrearIterador() Iterador {
	return &RadioIterador{radio: r}
}

func (r *Radio) Registrar(estacion string) {
	r.Estaciones = append(r.Estaciones, estacion)
}

// Iterador Concreto
type RadioIterador struct {
	radio  *Radio
	indice int
}

func (ri *RadioIterador) Valor() string {
	return ri.radio.Estaciones[ri.indice]
}

func (ri *RadioIterador) Siguiente() {
	ri.indice++
}

func (ri *RadioIterador) Anterior() {
	ri.indice--
}
