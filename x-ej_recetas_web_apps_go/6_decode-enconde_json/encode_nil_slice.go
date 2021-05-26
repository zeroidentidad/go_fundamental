package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var nilSlice []string
	emptySlice := []string{}

	m := map[string][]string{
		"nilSlice":   nilSlice,
		"emptySlice": emptySlice,
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()
}
