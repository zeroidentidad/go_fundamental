package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// escribir datos en un archivo
	fmt.Println("writing data into a file")
	writeFile("welcome to go")
	readFile()

	// leer datos de un archivo
	fmt.Println("reading data from a file")
	readFile()

}

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("./temp/testgo.txt", bytes, 0644)
	fmt.Println("created a file")
}
func readFile() {
	data, _ := ioutil.ReadFile("./temp/testgo.txt")
	fmt.Println("file content:")
	fmt.Println(string(data))
}
