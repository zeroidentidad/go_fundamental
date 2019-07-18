package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	nombre := "Jesus Ferrer"
	nombre2 := "Ana Karent"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Go Funca!</title>
		</head>
		<body>
		<h1>` + nombre + `</h1>
		<h2>` + nombre2 + `</h2>
		</body>
		</html>
			`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creaando archivo", err)
	}

	io.Copy(nf, strings.NewReader(str))
}
