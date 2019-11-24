package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("HOST", "localhost")

	env := os.Getenv("HOST")
	fmt.Println(env)
}
