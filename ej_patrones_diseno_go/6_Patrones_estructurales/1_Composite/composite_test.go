package composite

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	empleadoA := &DesarrolladorJunior{}
	empleadoB := &DesarrolladorJunior{}
	empleadoC := &DesarrolladorSenior{}

	gerenciaIT := &GerenciaIT{}
	gerenciaIT.AgregarEmpleado(empleadoA)
	gerenciaIT.AgregarEmpleado(empleadoB)
	gerenciaIT.AgregarEmpleado(empleadoC)

	fmt.Printf("El salario individual del desarrollador A es de $%d\n", empleadoA.ObtenerSalario())
	fmt.Printf("El salario individual del desarrollador B es de $%d\n", empleadoB.ObtenerSalario())
	fmt.Printf("El salario individual del desarrollador C es de $%d\n", empleadoC.ObtenerSalario())
	fmt.Printf("Los salarios de todos los desarrolladores de la Gerencia es de $%d\n", gerenciaIT.ObtenerSalario())
}
