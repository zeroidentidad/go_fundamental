package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func createURL() string {
	u, err := url.Parse("/params")

	if err != nil {
		panic(err)
	}

	u.Host = "localhost:8080"
	u.Scheme = "http"

	query := u.Query()
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()

	return u.String()
}

func main() {
	url := createURL()
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("encabezado", "valor")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	//fmt.Println("URL formada: ", url)

	fmt.Println("Header: ", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Body: ", string(body))
	fmt.Println("Status: ", response.Status)
}

//ejecutar 2do
