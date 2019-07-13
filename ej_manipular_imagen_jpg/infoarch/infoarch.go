package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../img/imagensk.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fi, _ := file.Stat() // Stat devuelve la estructura de FileInfo que describe el archivo.

	fmt.Println("Name: \t\t", fi.Name())
	fmt.Println("Size: \t\t", fi.Size())
	fmt.Println("Mode: \t\t", fi.Mode())
	fmt.Println("ModTime: \t", fi.ModTime())
	fmt.Println("IsDir: \t\t", fi.IsDir())
	fmt.Println("Sys: \t\t", fi.Sys())

}
