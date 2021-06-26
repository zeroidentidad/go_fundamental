package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.packtpub.com/product/hands-on-dependency-injection-in-go/9781789132762")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	stringBody := string(data)
	re := regexp.MustCompile(`.*main-book-price.*\n.*(\$[0-9]*\.[0-9]{0,2})`)

	priceMatches := re.FindStringSubmatch(stringBody)
	fmt.Printf("Book Price: %s\n", priceMatches[1])
}
