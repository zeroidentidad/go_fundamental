package modelos

// estructura
type CuentaCorriente struct {
	//...
}

// emular propiedad estática
// utilizando variable de paquete
var banco = "Banco S.A."

// emular método estático de CuentaCorriente
// utilizando una función de paquete
func CuentaCorrienteGetBanco() string {
	return banco
}
