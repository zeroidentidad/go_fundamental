package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("HOST", "localhost")

	//os.Unsetenv("HOST")

	env := os.Getenv("HOST")
	fmt.Println(env)
}
