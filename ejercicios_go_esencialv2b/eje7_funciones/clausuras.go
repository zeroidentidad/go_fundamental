package main

import "fmt"

func main() {
	fmt.Println(ganado(1))
}

func ganado(num int) (int, string) {

	vacas := func() int {
		return num * 10
	}

	return vacas() + 1, "vacas"
}
