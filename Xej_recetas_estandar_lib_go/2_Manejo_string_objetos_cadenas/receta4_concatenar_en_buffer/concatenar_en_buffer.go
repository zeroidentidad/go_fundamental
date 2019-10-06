package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"Esto ", "es ", "aun ",
		"mas ", "eficaz "}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}
