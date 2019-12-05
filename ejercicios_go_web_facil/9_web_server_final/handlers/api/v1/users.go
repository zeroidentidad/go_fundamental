package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../../../models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

func getUserByRequest(r *http.Request) (*models.User, error) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	user := models.GetUserById(userID)

	if user.ID == 0 {
		return user, errors.New("Usuario no existe en BD")
	}

	return user, nil
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.SetPassword(user.Password)
	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	models.SendData(w, user)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	request := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(request); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.Username = request.Username
	user.Email = request.Email
	user.SetPassword(request.Password)

	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		user.Delete()
		models.SendNoContent(w)
	}
}
