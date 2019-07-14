package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

type pixel struct {
	r, g, b, a uint32
}

func main() {

	images := getImages("../img/")

	// rango sobre el [] teniendo el []pixel, por ejemplo, de cada img
	// rango sobre el []pixel contiene los píxeles, por ejemplo, de cada píxel
	for i, img := range images {
		for j, pixel := range img {
			fmt.Println("Imagen", i, "\t pixel", j, "\t r g b a:", pixel)
			if j == 10 {
				break
			}
		}
	}
}

func getImages(dir string) [][]pixel {

	var images [][]pixel

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		img := loadImage(path)
		pixels := getPixels(img)
		images = append(images, pixels)
		return nil
	})

	return images
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

func getPixels(img image.Image) []pixel {

	limites := img.Bounds()
	fmt.Println(limites.Dx(), " x ", limites.Dy()) // debugging
	pixels := make([]pixel, limites.Dx()*limites.Dy())

	for i := 0; i < limites.Dx()*limites.Dy(); i++ {
		x := i % limites.Dx()
		y := i / limites.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
	}

	return pixels
}
