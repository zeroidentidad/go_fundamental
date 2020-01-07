package main

import "fmt"

func main() {
	var i, j = 0, 0
	for i < 10 {
		fmt.Printf("\n valor de i: %d", i)
		if i == 5 {
			fmt.Printf("\t\t multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Printf("\t\t paso por aqui \n")
		i++
	}

RUTINA:
	for j < 10 {
		if j == 4 {
			j = j + 2
			fmt.Println("\n voy a RUTINA sumando 2 a 1")
			goto RUTINA
		}
		fmt.Printf("\n valor de j: %d \n", j)
		j++
	}
}
