package main

import (
	"fmt"
	"os"
)

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("./test/testforread.txt")

	checkerror(err)

	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	b1 := make([]byte, fileinfo.Size())

	f.Read(b1)
	checkerror(err)
	fmt.Printf(string(b1))
	fmt.Println()

	f.Close()
}
