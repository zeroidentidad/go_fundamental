package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "No habites en el pasado, no sueñes con el futuro, concentra la mente en el presente. -Buddha "
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	if err != nil {
		fmt.Println(err)
		return
	}

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Get("http://www.notexists.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		fmt.Println(err)
	}

}
