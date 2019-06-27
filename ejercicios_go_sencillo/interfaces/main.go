package main

import (
	"bytes"
	"fmt"

	"./interfaces"
)

func main() {
	in := bytes.NewReader([]byte("ejemplo"))
	out := &bytes.Buffer{}
	fmt.Print("stdout en Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout en PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
