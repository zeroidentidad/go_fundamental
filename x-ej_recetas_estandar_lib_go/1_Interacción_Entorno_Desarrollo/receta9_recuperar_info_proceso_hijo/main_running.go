package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep" // linux / mac
	}
	proc := exec.Command(cmd, "1")
	proc.Start()

	// No se devuelve ningún estado de proceso
	// hasta que finalice el proceso.
	fmt.Printf("Estado para ejecutar el proceso: %v\n", proc.ProcessState)

	// El PID se puede obtener
	// del evento para el proceso en ejecución
	fmt.Printf("PID del proceso en ejecución: %d\n\n", proc.Process.Pid)
}
