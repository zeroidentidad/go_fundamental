package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)

	buf := bytes.NewBuffer([]byte{})

	// El buffer que implementa
	// la interfaz io.Writer se asigna a
	// Stdout
	proc.Stdout = buf

	// Para evitar condiciones de carrereo
	// en este ejemplo. Esperamos hasta
	// la salida del proceso.
	proc.Run()

	// El proceso escribe la salida en
	// en búfer y usamos los bytes
	// para imprimir la salida.
	fmt.Println(string(buf.Bytes()))

}
