package main

import (
	"log"
	"net/http"

	"encoding/json"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ke pedo"))
}

type Response struct {
	Mensaje string `json:"mensaje"`
	Valid   bool   `json:"valid"`
}

func HolaJSON(w http.ResponseWriter, r *http.Request) {
	response := CreateResponse("Ke onda es JSON", true)
	json.NewEncoder(w).Encode(response)
}

func CreateResponse(mensaje string, valid bool) Response {
	return Response{
		mensaje,
		valid,
	}
}

func Html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/index.html")
}

type User struct {
	UserName  string
	Websocket *websocket.Conn
}

var Users = struct {
	m map[string]User
	sync.RWMutex
}{m: make(map[string]User)}

func ValidarUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")

	response := Response{}
	if UserExist(username) {
		//no permitir otro ingreso
		response.Mensaje = "No es valido"
		response.Valid = false
	} else {
		//permitir ingreso
		response.Mensaje = "Es valido"
		response.Valid = true
	}
	json.NewEncoder(w).Encode(response)
}

func UserExist(username string) bool {
	Users.RLock() //lectura
	defer Users.RUnlock()

	if _, ok := Users.m[username]; ok {
		return true
	}
	return false
}

func CreateUser(username string, ws *websocket.Conn) User {
	return User{
		username,
		ws,
	}
}

func AddUser(user User) {
	Users.Lock() //escritura
	defer Users.Unlock()
	Users.m[user.UserName] = user
}

func RemoveUser(username string) {
	Users.Lock() //escritura
	defer Users.Unlock()
	delete(Users.m, username) //borrar del map
}

func SendMessage(typeMsg int, msg []byte) {
	Users.RLock() //lectura
	defer Users.RUnlock()

	for _, user := range Users.m {
		if err := user.Websocket.WriteMessage(typeMsg, msg); err != nil {
			return
		}
	}
}

func ArrayByte(value string) []byte {
	return []byte(value)
}

func ConcatMsg(username string, arr []byte) string {
	return username + ": " + string(arr[:]) //todo
}

func WebSocket(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if err != nil {
		log.Println("Error:", err)
		return
	}

	currentUser := CreateUser(username, ws)
	AddUser(currentUser)
	log.Println("Usuario agregado")

	for {
		typeMsg, msg, err := ws.ReadMessage()

		if err != nil {
			RemoveUser(username)
			return
		}

		finalMsg := ConcatMsg(username, msg)
		SendMessage(typeMsg, ArrayByte(finalMsg))
	}
}

func main() {

	cssHandle := http.FileServer(http.Dir("./frontend/css/"))
	jsHandle := http.FileServer(http.Dir("./frontend/js/"))

	mux := mux.NewRouter()
	mux.HandleFunc("/", Hola).Methods("GET")
	mux.HandleFunc("/json", HolaJSON).Methods("GET")
	mux.HandleFunc("/html", Html).Methods("GET")
	mux.HandleFunc("/validar", ValidarUser).Methods("POST")
	mux.HandleFunc("/chat/{username}", WebSocket).Methods("GET")

	http.Handle("/", mux) //se traducen a las rutas de mux
	http.Handle("/css/", http.StripPrefix("/css/", cssHandle))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandle))

	log.Println("Server ejecutandose")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
