package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Printf(
			"usage: %s add|del \n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		fmt.Println("adding item")
	case "del":
		fmt.Println("deleting item")
	default:
		fmt.Printf(
			"usage: %s add|del \n", os.Args[0])
		os.Exit(1)
	}

}
