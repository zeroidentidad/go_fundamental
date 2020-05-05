package abstractfactory

// Producto Abstracto Interface
type Puerta interface {
	VerMaterial() string
}

// Producto Concreto
type PuertaMadera struct{}

func (pm *PuertaMadera) VerMaterial() string {
	return "Madera"
}

// Producto Concreto
type PuertaMetal struct{}

func (pm *PuertaMetal) VerMaterial() string {
	return "Metal"
}

// Fábrica Abstracta Interface
type FabricaPuerta interface {
	ConstruirPuerta() Puerta
}

// Fábrica Concreta
type FabricaPuertaMadera struct{}

func (fpm *FabricaPuertaMadera) ConstruirPuerta() Puerta {
	return &PuertaMadera{}
}

// Fábrica Concreta
type FabricaPuertaMetal struct{}

func (fpm *FabricaPuertaMetal) ConstruirPuerta() Puerta {
	return &PuertaMetal{}
}
