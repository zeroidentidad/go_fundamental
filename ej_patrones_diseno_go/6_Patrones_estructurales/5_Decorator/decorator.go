package decorator

// Componente Interface
type CafeInterface interface {
	GetCosto() int
	GetDetalle() string
}

// Componente Concreto
type Cafe struct{}

func (c *Cafe) GetCosto() int {
	return 30
}

func (c *Cafe) GetDetalle() string {
	return "Cafe"
}

// Decorador
type CafeDecorador struct {
	CafeInterface
}

func (cd *CafeDecorador) GetCosto() int {
	return cd.CafeInterface.GetCosto()
}

func (cd *CafeDecorador) GetDetalle() string {
	return cd.CafeInterface.GetDetalle()
}

// Decorador Concreto
type CafeConCrema struct {
	*CafeDecorador
}

func (ccc *CafeConCrema) GetCosto() int {
	return ccc.CafeDecorador.GetCosto() + 15
}

func (ccc *CafeConCrema) GetDetalle() string {
	return ccc.CafeDecorador.GetDetalle() + " con crema"
}

// Decorador Concreto
type CafeConCanela struct {
	*CafeDecorador
}

func (ccc *CafeConCanela) GetCosto() int {
	return ccc.CafeDecorador.GetCosto() + 10
}

func (ccc *CafeConCanela) GetDetalle() string {
	return ccc.CafeDecorador.GetDetalle() + " con canela"
}
