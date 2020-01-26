package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// El escáner puede escanear entradas por líneas
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Entrada: %s\n", txt)
	}

}
