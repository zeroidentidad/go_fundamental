package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue // omitir aqui y volver para continuar
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
