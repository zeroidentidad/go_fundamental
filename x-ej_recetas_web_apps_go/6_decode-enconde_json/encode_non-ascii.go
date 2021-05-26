package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	s := struct {
		CostUSD string `json:"cost $"` // Esto esta bien.
		CostEUR string `json:"cost €"` // Contiene carácter no ASCII € y se ignorará.
	}{
		CostUSD: "100.00",
		CostEUR: "100.00",
	}

	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()

}
