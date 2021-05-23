package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	_json := `[1, 2, 3]`

	var shortArray [2]int
	err := json.NewDecoder(strings.NewReader(_json)).Decode(&shortArray)
	if err != nil {
		log.Fatal(err)
	}

	var longArray [4]int
	err = json.NewDecoder(strings.NewReader(_json)).Decode(&longArray)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("shortArray: %v\n", shortArray)
	fmt.Printf("longArray: %v\n", longArray)
}
