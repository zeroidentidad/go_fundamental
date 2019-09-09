package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	// parametros al proceso
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	// Se obtiene la entrada del proceso
	// en forma de io.WriteCloser. La
	// implementación usa el os.Pipe
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// Para fines de depuración vemos la
	// salida del proceso ejecutadose
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Programa dice: " + s.Text())
		}
	}()

	// Iniciar el proceso
	proc.Start()

	// Ahora las siguientes líneas
	// se escriben al proceso
	// hijo de entrada estándar
	fmt.Println("Escribiendo entrada.")
	io.WriteString(stdin, "Hola\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "es genial\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()

}
