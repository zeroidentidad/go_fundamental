package main

import (
	"crypto"
)

func main() {
	//aunque es sintácticamente correcto
	// logicamente ocasiona panic error
	crypto.SHA1.New()
}
