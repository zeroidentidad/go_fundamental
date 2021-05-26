package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	s := []string{
		"<foo>",
		"bar & baz",
	}

	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()
}
