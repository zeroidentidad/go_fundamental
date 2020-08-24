package database

import (
	"context"
	"time"
)

// ExecWithTimeout terminará el intento
// para obtener la hora actual
func ExecWithTimeout() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// terminar el tiempo de espera inmediatamente
	ctx, can := context.WithDeadline(ctx, time.Now())

	// llamar cancelar después de completar
	defer can()

	// la transacción es consciente del contexto
	_, err = db.BeginTx(ctx, nil)
	return err
}
