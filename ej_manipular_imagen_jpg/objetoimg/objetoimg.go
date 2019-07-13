package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	img := loadImage("../img/imagensk.jpg")
	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Printf("%d %d %d %d \n", r, g, b, a) // 13155 12323 7996 65535 //tonalidad de colores dominantes
}

func loadImage(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
