package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"../models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserXML{ID: 1, Username: "Zero", Password: "xd"}
	w.Header().Set("Content-Type", "text/xml")
	userxml, _ := xml.Marshal(user)
	fmt.Fprintf(w, string(userxml))
}
