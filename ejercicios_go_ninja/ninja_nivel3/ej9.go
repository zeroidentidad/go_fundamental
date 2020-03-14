package main

import (
	"fmt"
)

func main() {
	deporteFav := "béisbol"
	switch deporteFav {
	case "béisbol":
		fmt.Println("Ve al estadio")
	case "natación":
		fmt.Println("Ve a la piscina")
	case "crossfit":
		fmt.Println("Te quiero ver en los crossfit games.")
	}
}
