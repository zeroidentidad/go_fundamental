package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("======================================================================")
	fmt.Println(":D Saludos terricola, ojala y te sirva mi repo para aprender sobre Go")
	fmt.Println("======================================================================")

	if runtime.GOOS == "windows" {
		prc := exec.Command("dir", "")
		out := bytes.NewBuffer([]byte{})
		prc.Stdout = out
		err := prc.Run()
		if err != nil {
			fmt.Println(err)
		}

		if prc.ProcessState.Success() {
			fmt.Println("Info sobre:")
			fmt.Println(out.String())
		}
	}

	if runtime.GOOS == "linux" {
		prc := exec.Command("ls", "-a")
		out := bytes.NewBuffer([]byte{})
		prc.Stdout = out
		err := prc.Run()
		if err != nil {
			fmt.Println(err)
		}

		if prc.ProcessState.Success() {
			fmt.Println("Info sobre:")
			fmt.Println(out.String())
		}
	}

}
