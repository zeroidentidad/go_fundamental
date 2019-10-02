package osfunc

import (
	"os/exec"
)

func CommandOutput(command string) []byte {
	b, _ := Setcommand(command).Output()
	var lenComand = len(b)
	if lenComand > 1 {
		return b[:lenComand-1]
	}
	return b
}

func CommandRun(command string) error {
	return Setcommand(command).Run()
}

func Setcommand(command string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", command)
}
