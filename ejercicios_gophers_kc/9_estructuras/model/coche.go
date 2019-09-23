package model

// Para declarar una struct:
// type <Nombre de la struct> struct { }

// para poner las variables que la componen es necesario
// tanto declarar con nombre como el tipo de la misma
//  <nombre> <tipo>

type Motor struct {
	NumCilindros int
	Cilindrada   int
}

type Coche struct {
	Motor         *Motor
	NumeroRuedas  int
	numeroDeSerie string
}
