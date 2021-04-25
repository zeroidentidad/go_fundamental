package main

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := "ping"
	timeout := 2 * time.Second

	// La herramienta de cmd
	// "ping" es ejecutado por
	// 2 segundos
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, cmd, "google.com")

	// Se obtiene la salida del proceso
	// en forma de io.ReadCloser. La
	// implementación usa el os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// Iniciar el proceso
	proc.Start()

	// Para una lectura más cómoda
	// se utiliza bufio.Scanner.
	// La llamada de lectura está bloqueando.
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
