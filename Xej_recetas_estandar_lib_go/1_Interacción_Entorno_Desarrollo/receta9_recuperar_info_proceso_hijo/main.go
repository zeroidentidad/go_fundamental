package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()

	// La función de espera,
	// espera hasta que termine el proceso.
	proc.Wait()

	// Después de que el proceso termina
	// el *os.ProcessState contiene
	// información simple
	// sobre el proceso ejecutado
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Proceso tardo: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Salida exitosa: %t\n", proc.ProcessState.Success())
}
