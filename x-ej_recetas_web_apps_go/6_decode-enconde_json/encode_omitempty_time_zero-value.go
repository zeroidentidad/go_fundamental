package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	s := struct {
		Foo time.Time `json:",omitempty"`
	}{}

	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()

}
