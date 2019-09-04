package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	prc := exec.Command("go", "run run1.go") //ls -a
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Proceso ejecutado con éxito con salida:")
		fmt.Println(out.String())
	}
}
