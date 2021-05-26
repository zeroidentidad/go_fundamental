package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	s := []float64{
		123.0,
		456.100,
		789.990,
	}

	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()
}
