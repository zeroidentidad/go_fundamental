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

	f, err := os.OpenFile("./test/testforread.txt", os.O_RDWR|os.O_CREATE, 0666)

	checkerror(err)

	fileinfo, err1 := f.Stat()
	checkerror(err1)

	f.Seek(fileinfo.Size(), 0)
	f.Write([]byte(" test watahea"))
	f.Sync()
	f.Seek(0, 0)

	fileinfo1, err2 := f.Stat()
	checkerror(err2)

	b1 := make([]byte, fileinfo1.Size())

	_, err3 := f.Read(b1)
	checkerror(err3)

	fmt.Printf(string(b1))
	fmt.Println()

	f.Close()
}
