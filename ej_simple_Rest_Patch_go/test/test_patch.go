package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func main() {
	url := "http://localhost:3000/user/actualizar/1"

	var jsonStr = []byte(fmt.Sprintf(`{"nombre":"NomPatch%v", "apellido":"ApePatch%v","edad": %v}`, rand.Intn(100), rand.Intn(100), rand.Intn(100)))
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))

	fmt.Println("URL:>", url)

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
