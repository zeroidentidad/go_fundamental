package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	mx, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		panic(err)
	}
	t := time.Date(2020, 11, 20, 11, 20, 10, 0, mx)

	// json.Marshaler interface
	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("Serializado como RFC 3339:", string(b))
	t2 := time.Time{}
	t2.UnmarshalJSON(b)
	fmt.Println("Deserializado de RFC 3339:", t2)

	// Serialize como epoca
	epoch := t.Unix()
	fmt.Println("Serializado como Epoca:", epoch)

	// Deserialize epoca
	jsonStr := fmt.Sprintf("{ \"created\":%d }", epoch)
	data := struct {
		Created int64 `json:"created"`
	}{}
	json.Unmarshal([]byte(jsonStr), &data)
	deserialized := time.Unix(data.Created, 0)
	fmt.Println("Deserializado de Epoca:", deserialized)

}
