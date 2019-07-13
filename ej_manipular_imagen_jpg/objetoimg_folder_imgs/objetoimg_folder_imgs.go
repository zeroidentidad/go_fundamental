package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var images []image.Image

	filepath.Walk("../img/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		img := loadImage(path)
		images = append(images, img)
		return nil
	})

	for _, v := range images {
		r, g, b, a := v.At(0, 0).RGBA()
		fmt.Printf("%d %d %d %d \n", r, g, b, a) // 13155 12323 7996 65535 //tonalidad de colores dominantes
	}

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
