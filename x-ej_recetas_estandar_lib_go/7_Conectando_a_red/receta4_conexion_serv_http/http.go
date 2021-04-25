package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hola mundo"),
	}
}

const addr = "localhost:8080"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()
	time.Sleep(time.Second * 5)

	useRequest()
	simplePost()
}

func simplePost() {
	res, err := http.Post("http://localhost:8080",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Nani&surname=Mazaka"))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Response from server:" + string(data))
}

func useRequest() {

	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "Zero")
	form.Add("surname", "Identidad")

	req, err := http.NewRequest("POST",
		"http://localhost:8080",
		strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded")

	res, err := hc.Do(req)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Response from server:" + string(data))
}
