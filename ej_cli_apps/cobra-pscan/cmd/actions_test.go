package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"

	"pscan/scan"
)

func TestHostActions(t *testing.T) {
	// Definir hosts para prueba de acciones
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	// Casos de prueba para prueba de acciones
	testCases := []struct {
		name           string
		args           []string
		expectedOut    string
		initList       bool
		actionFunction func(io.Writer, string, []string) error
	}{
		{
			name:           "AddAction",
			args:           hosts,
			expectedOut:    "Added host: host1\nAdded host: host2\nAdded host: host3\n",
			initList:       false,
			actionFunction: addAction,
		},
		{
			name:           "ListAction",
			expectedOut:    "host1\nhost2\nhost3\n",
			initList:       true,
			actionFunction: listAction,
		},
		{
			name:           "DeleteAction",
			args:           []string{"host1", "host2"},
			expectedOut:    "Deleted host: host1\nDeleted host: host2\n",
			initList:       true,
			actionFunction: deleteAction,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup prueba de acción
			tf, cleanup := setup(t, hosts, tc.initList)
			defer cleanup()

			// Definir var para capturar salida de acción
			var out bytes.Buffer

			// Ejecutar acción y capturar salida
			if err := tc.actionFunction(&out, tf, tc.args); err != nil {
				t.Fatalf("Expected no error, got %q\n", err)
			}

			// Salida de acciones de prueba
			if out.String() != tc.expectedOut {
				t.Errorf("Expected output %q, got %q\n", tc.expectedOut, out.String())
			}
		})
	}
}

func TestScanAction(t *testing.T) {
	// Definir hosts para prueba de escaneo
	hosts := []string{
		"localhost",
		"unknownhostoutthere",
	}

	// Setup prueba de escaneo
	tf, cleanup := setup(t, hosts, true)
	defer cleanup()

	ports := []int{}

	// Init ports, 1 open, 1 closed
	for i := 0; i < 2; i++ {
		ln, err := net.Listen("tcp", net.JoinHostPort("localhost", "0"))
		if err != nil {
			t.Fatal(err)
		}

		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal(err)
		}

		ports = append(ports, port)

		if i == 1 {
			ln.Close()
		}
	}

	// Definir salida esperada para la acción de escaneo
	expectedOut := fmt.Sprintln("localhost:")
	expectedOut += fmt.Sprintf("\t%d: open\n", ports[0])
	expectedOut += fmt.Sprintf("\t%d: closed\n", ports[1])
	expectedOut += fmt.Sprintln()
	expectedOut += fmt.Sprintln("unknownhostoutthere: Host not found")
	expectedOut += fmt.Sprintln()

	// Definir var para capturar salida de escaneo
	var out bytes.Buffer

	// Ejecutar escaneo y capturar salida
	if err := scanAction(&out, tf, ports); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	// Salida de escaneo de prueba
	if out.String() != expectedOut {
		t.Errorf("Expected output %q, got %q\n", expectedOut, out.String())
	}
}

func TestIntegration(t *testing.T) {
	// Definir hosts para prueba de integración
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	// Setup prueba de integración
	tf, cleanup := setup(t, hosts, false)
	defer cleanup()

	delHost := "host2"

	hostsEnd := []string{
		"host1",
		"host3",
	}

	// Definir var para capturar salida
	var out bytes.Buffer

	// Definir resultado esperado para todas las acciones
	expectedOut := ""
	for _, v := range hosts {
		expectedOut += fmt.Sprintf("Added host: %s\n", v)
	}
	expectedOut += strings.Join(hosts, "\n")
	expectedOut += fmt.Sprintln()
	expectedOut += fmt.Sprintf("Deleted host: %s\n", delHost)
	expectedOut += strings.Join(hostsEnd, "\n")
	expectedOut += fmt.Sprintln()
	for _, v := range hostsEnd {
		expectedOut += fmt.Sprintf("%s: Host not found\n", v)
		expectedOut += fmt.Sprintln()
	}

	// Agregar hosts a la lista
	if err := addAction(&out, tf, hosts); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	// Listar hosts
	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	// Eliminar host2
	if err := deleteAction(&out, tf, []string{delHost}); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	// Listar hosts después de eliminar
	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	// Escanear hosts
	if err := scanAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	// Salida de prueba de integración
	if out.String() != expectedOut {
		t.Errorf("Expected output %q, got %q\n", expectedOut, out.String())
	}
}

func setup(t *testing.T, hosts []string, initList bool) (string, func()) {
	// Crear archivo temporal
	tf, err := ioutil.TempFile("", "pScan")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	// Inicializar lista de ser necesaria
	if initList {
		hl := &scan.HostsList{}

		for _, h := range hosts {
			hl.Add(h)
		}

		if err := hl.Save(tf.Name()); err != nil {
			t.Fatal(err)
		}
	}

	// Retornar nombre del archivo temporal y función de cleanup
	return tf.Name(), func() {
		os.Remove(tf.Name())
	}
}
