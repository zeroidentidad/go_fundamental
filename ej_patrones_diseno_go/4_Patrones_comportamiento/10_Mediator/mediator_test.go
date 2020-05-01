package mediator

import "testing"

// Test de uso del patrón
func TestPattern(t *testing.T) {
	mediador := &ChatRoom{}

	usuarioA := &UsuarioChat{"Daniel", mediador}
	usuarioB := &UsuarioChat{"Pedro", mediador}

	usuarioA.EnviarMensaje("Hola como estas?")
	usuarioB.EnviarMensaje("Muy bien y vos?")
}
