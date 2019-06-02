package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../commons"
	"../configuration"
	"../models"
)

//CommentCreate maneja la creación de comentarios
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error en lectura de comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error en inserción de comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado"
	commons.DisplayMessage(w, m)
}
