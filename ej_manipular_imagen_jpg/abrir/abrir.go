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
	defer file.Close()

	fmt.Println(file)
}
