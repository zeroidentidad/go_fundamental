package main

import (
	"fmt"

	"./basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("error personalizado: ", err)
}
