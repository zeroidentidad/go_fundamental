package composite

// Componente Interface
type Empleado interface {
	ObtenerSalario() int
}

// Hoja
type DesarrolladorSenior struct{}

func (ds *DesarrolladorSenior) ObtenerSalario() int {
	return 1000
}

// Hoja
type DesarrolladorJunior struct{}

func (dj *DesarrolladorJunior) ObtenerSalario() int {
	return 750
}

// Compuesto
type GerenciaIT struct {
	empleados []Empleado
}

func (g *GerenciaIT) AgregarEmpleado(empleado Empleado) {
	g.empleados = append(g.empleados, empleado)
}

func (g *GerenciaIT) ObtenerSalario() int {
	sumaSalarios := 0

	for _, empleado := range g.empleados {
		sumaSalarios = sumaSalarios + empleado.ObtenerSalario()
	}

	return sumaSalarios
}
