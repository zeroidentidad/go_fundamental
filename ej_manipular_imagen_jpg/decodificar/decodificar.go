package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	file, err := os.Open("../img/imagensk.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TIPO IMAGEN: %T \n", img)

	limites := img.Bounds()

	fmt.Println("Alto x Ancho: ", limites.Dx(), " x ", limites.Dy())
	fmt.Println("Total Pixeles: ", limites.Dx()*limites.Dy())

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Println("Primer pixel:", r, g, b, a)

	r, g, b, a = img.At(limites.Dx(), limites.Dy()).RGBA()
	fmt.Println("Ultimo pixel:", r, g, b, a)
}
