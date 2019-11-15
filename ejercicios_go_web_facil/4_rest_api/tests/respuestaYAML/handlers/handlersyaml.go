package handlers

import (
	"fmt"
	"net/http"

	"../models"
	"gopkg.in/yaml.v2"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserYAML{ID: 1, Username: "Zero", Password: "xd"}
	//w.Header().Set("Content-Type", "text/xml")
	useryaml, _ := yaml.Marshal(user)
	fmt.Fprintf(w, string(useryaml))
}

// go get gopkg.in/yaml.v2
