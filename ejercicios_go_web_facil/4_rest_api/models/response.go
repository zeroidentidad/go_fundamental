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

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}
}

func (this *Response) Send() {
	this.writer.WriteHeader(this.Status)
	this.writer.Header().Set("Content-Type", this.contentType)

	responsejson, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(responsejson))
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFoundResponse("Recurso no disponible")
	response.Send()
}

func (this *Response) NotFoundResponse(msg string) {
	this.Status = http.StatusNotFound
	this.Message = msg
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity("Entidad no procesable")
	response.Send()
}

func (this *Response) UnprocessableEntity(msg string) {
	this.Status = http.StatusUnprocessableEntity
	this.Message = msg
}
