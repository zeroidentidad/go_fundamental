package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	user, _ := reader.ReadString('\n')

	fmt.Print("Enter password: ")
	pass, _ := term.ReadPassword(int(syscall.Stdin))

	fmt.Printf("\n\nUsername: %sPassword: %s\n", user, string(pass))
}
