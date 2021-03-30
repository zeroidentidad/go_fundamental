package unit

import (
	"pastein/cmd/web"
	"testing"
	"time"
)

// go test -run TestHumanDate -v
// -failfast flag para detener ejecución pruebas después del primer error
// go test -run TestHumanDate -v -failfast
func TestHumanDate(t *testing.T) {
	// table-driven tests:
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2021, 12, 17, 10, 0, 0, 0, time.UTC),
			want: "17 Dec 2021 at 10:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2021, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Dec 2021 at 09:00",
		},
	}

	// Loop test cases.
	for _, tt := range tests {
		// run sub-test
		t.Run(tt.name, func(t *testing.T) {
			// Inicializar un nuevo objeto time.Time y pasarlo a función humanDate.
			hd := web.TestHumanDate(tt.tm)
			/*Comprobar que la salida de la función humanDate esté en el formato que esperamos.
			  Si no es lo que esperamos, usar función t.Errorf() para indicar que la prueba
			  ha fallado y log de los valores esperados y reales. */
			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}
