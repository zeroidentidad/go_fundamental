package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	s := struct {
		Foo struct {
			Bar string `json:",omitempty"`
		} `json:",omitempty"`
	}{}

	s2 := struct {
		Foo *struct {
			Bar string `json:",omitempty"`
		} `json:",omitempty"`
	}{}

	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}

	b2, err := json.Marshal(s2)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()

	os.Stdout.Write(b2)
	fmt.Println()
}
