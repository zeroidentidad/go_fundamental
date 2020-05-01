package observer

import "fmt"

// Interface Sujeto
type Sujeto interface {
	Adquirir(Observador)
	notificar()
}

// Sujeto Concreto
type AgenciaEmpleo struct {
	observadores []Observador
}

func (ae *AgenciaEmpleo) Adquirir(observador Observador) {
	ae.observadores = append(ae.observadores, observador)
}

func (ae *AgenciaEmpleo) notificar(oferta string) {
	for _, observador := range ae.observadores {
		observador.Actualizar(oferta)
	}
}

func (ae *AgenciaEmpleo) IngresarOfertaLaboral(oferta string) {
	ae.notificar(oferta)
}

// Interface Observador
type Observador interface {
	Actualizar(string)
}

// Observador Concreto
type ObservadorEmpleo struct {
	nombre string
}

func (o *ObservadorEmpleo) Actualizar(oferta string) {
	fmt.Printf("Hola %s, existe la siguiente oferta de empleo: %s\n", o.nombre, oferta)
}
