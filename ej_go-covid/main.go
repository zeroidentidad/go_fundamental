package main

import (
	"fmt"

	"github.com/ahmednafies/go-covid/covid"
)

func main() {
	data := covid.GetCountryByName("mexico")
	fmt.Println(data.Attrs.Id)
	fmt.Println(data.Attrs.Country)
	fmt.Println(data.Attrs.LastUpdate)
	fmt.Println(data.Attrs.Confirmed)
	fmt.Println(data.Attrs.Deaths)
	fmt.Println(data.Attrs.Active)
	fmt.Println(data.Attrs.Recovered)
	fmt.Println(data.Attrs.Latitude)
	fmt.Println(data.Attrs.Longitude)
}
