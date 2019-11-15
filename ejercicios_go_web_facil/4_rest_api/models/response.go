package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

func GetDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}
}

func (this *Response) NotFoundResponse(msg string) {
	this.Status = http.StatusNotFound
	this.Data = nil
	this.Message = msg
}

func (this *Response) Send() {
	this.writer.WriteHeader(this.Status)
	this.writer.Header().Set("Content-Type", this.contentType)

	responsejson, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(responsejson))
}
