package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("sample.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go es genial :v !\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Si! Golang es genial.\n"))
	if err != nil {
		panic(err)
	}
}
