package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type teststruct struct {
	Val1 string
	Val2 int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("testheader", "hval1")
		//fmt.Fprintf(w, prepresponse(r.FormValue("tst")))
		var ts teststruct
		ts.Val1 = "test value 1"
		ts.Val2 = 2
		//w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, prepresponse(ts))
	})

	http.ListenAndServe(":8080", nil)
}

func prepresponse(obj interface{}) string {
	println(obj)
	resp, _ := json.Marshal(obj)

	println(string(resp))
	return string(resp)
}
