package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../img/imagensk.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs, _ := ioutil.ReadAll(file)
	fmt.Println(bs)
}
