package contextlog

import (
	"context"

	"github.com/apex/log"
)

type key int

// logFields es una clave a usar
// para nuestro registro de contexto
const logFields key = 0

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// FromContext toma una entrada y un contexto
// luego devuelve una entrada poblada del objeto de contexto
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField agrega un campo de registro al contexto
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}

// WithFields agrega muchos campos de registro al contexto
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	ctx = context.WithValue(ctx, logFields, f)
	return ctx
}
