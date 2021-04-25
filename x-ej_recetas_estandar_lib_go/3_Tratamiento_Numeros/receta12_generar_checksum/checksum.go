package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var content = "Esto es contenido para verificar"

func main() {

	checksum := MD5(content)
	checksum2 := FileMD5("content.dat")

	fmt.Printf("Checksum 1: %s\n", checksum)
	fmt.Printf("Checksum 2: %s\n", checksum2)
	if checksum == checksum2 {
		fmt.Println("Suma de contenido coincide")
	}

}

// MD5 crea el hash md5 para contenido dado codificado en
// cadena hexadecimal
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// FileMD5 crea hash md5 codificado hexadecimal de contenido de archivo
func FileMD5(path string) string {

	h := md5.New()
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(h, f)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
