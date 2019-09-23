package model

import (
	"fmt"
)

// Para declarar una struct:
// type <Nombre de la struct> struct { }

// para poner las variables que la componen es necesario
// tanto declarar con nombre como el tipo de la misma
//  <nombre> <tipo>
// modelo privado
type motor struct {
	NumCilindros int
	Cilindrada   int
}

// model publico

type Coche struct {
	Motor         *motor
	NumeroRuedas  int
	numeroDeSerie string
	roto          bool
}

// En las estructuras siempre es bueno trabajar con punteros

func NewCoche(cilindros, cilindrada, numRuedas int) *Coche {
	var motorCoche *motor
	if cilindros > 0 {
		motorCoche = &motor{
			NumCilindros: cilindros,
			Cilindrada:   cilindrada,
		}
	}
	return &Coche{
		Motor:         motorCoche,
		NumeroRuedas:  numRuedas,
		numeroDeSerie: "Adwfsd",
		roto:          false,
	}
}

func (c *Coche) Arranca() string {
	if c.roto {
		return "No ha podido arrancar, por favor llevame al mecanico"
	}
	return fmt.Sprintf("El motor ha arrancado correctamente con %v de cilindrada", c.Motor.Cilindrada)
}

func (c *Coche) IncrementarPotencia(incremento int) {
	c.Motor.Cilindrada += incremento
}

func (c *Coche) PincharRueda() {
	if c.NumeroRuedas == 1 {
		c.NumeroRuedas = 0
		c.roto = true
	} else {
		c.NumeroRuedas--
	}

}
