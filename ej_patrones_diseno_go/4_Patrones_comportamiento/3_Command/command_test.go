package command

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	invocador := Invocador{comandos: []Comando{}}
	receptor := Receptor{}

	// se establecen dos comandos concretos y se los elimina
	// el invocador queda sin comandos que ejecutar
	invocador.GuardarComando(ComandoPrender{receptor: receptor})
	invocador.GuardarComando(ComandoApagar{receptor: receptor})
	invocador.Limpiar()

	// se establecen dos comandos concretos iguales y se elimina el último
	// el invocador queda con un único comando para ejecutar
	invocador.GuardarComando(ComandoPrender{receptor: receptor})
	invocador.GuardarComando(ComandoPrender{receptor: receptor})
	invocador.EliminarUltimoComando()

	// se establece un comando concreto más
	invocador.GuardarComando(ComandoApagar{receptor: receptor})

	// el invocador ejecuta los dos comandos
	fmt.Printf("Comandos ejecutados:\n%v", invocador.Ejecutar())
}
