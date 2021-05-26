package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	m := map[int]string{
		123:     "foo",
		456_000: "bar",
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()

	t1 := time.Now()
	t2 := t1.Add(24 * time.Hour)
	m2 := map[time.Time]string{
		t1: "foo",
		t2: "bar",
	}

	b2, err := json.Marshal(m2)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b2)
	fmt.Println()
}
