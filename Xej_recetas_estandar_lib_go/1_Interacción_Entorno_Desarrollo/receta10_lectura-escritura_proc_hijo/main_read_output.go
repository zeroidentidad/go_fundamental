package main

import (
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

	// La salida ejecutará el proceso
	// termina y devuelve la salida
	// estandar en un segmento de byte.
	buff, err := proc.Output()

	if err != nil {
		panic(err)
	}

	// La salida del proceso hijo
	// en forma de slice de bytes
	// impreso como cadena
	fmt.Println(string(buff))

}
