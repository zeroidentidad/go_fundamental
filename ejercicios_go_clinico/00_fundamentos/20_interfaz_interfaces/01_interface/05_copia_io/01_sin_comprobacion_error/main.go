package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "No habites en el pasado, no sueñes con el futuro, concentra la mente en el presente. -Buddha "
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("http://www.mcleods.com")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
