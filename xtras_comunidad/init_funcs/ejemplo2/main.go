package main

import (
	"fmt"

	_ "ejemplo2/a"
)

func init() {
	fmt.Println("Init desde mi programa")
}

func main() {
	fmt.Println("Mi programa")
}
