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

	r, g, b, a := img.At(0, 0).RGBA()
	fmt.Printf("%T %T %T %T \n", r, g, b, a) // uint32 uint32 uint32 uint32
	fmt.Printf("%v %v %v %v \n", r, g, b, a) // 13155 12323 7996 65535
}
