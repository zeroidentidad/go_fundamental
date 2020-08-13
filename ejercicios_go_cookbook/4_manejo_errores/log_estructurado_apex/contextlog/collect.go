package contextlog

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Inicializa las llamadas a 3 funciones para configurar,
// luego registrar antes de terminar
func Initialize() {
	// configurar registro básico
	log.SetHandler(text.New(os.Stdout))
	// inicializar contexto
	ctx := context.Background()
	// crear un registrador y vincularlo al contexto
	ctx, e := FromContext(ctx, log.Log)

	// establecer un campo
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Something")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"state": "Tabasco", "country": "MX"})
}
