package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/olahol/melody"
)

var mel *melody.Melody
var usuarios map[string]string

const token = "8e47f8c9-858e-4c9b-8bc6-9b2aa4bff448"

type ResponseWebSocket struct {
	Tipo         string `json:"type"`
	DataResponse `json:"data_response"`
}
type DataResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ResponseSale struct {
	Tipo     string `json:"type"`
	DataSale `json:"data_sale"`
}

type DataSale struct {
	Product  float64 `json:"product"`
	Quantity float64 `json:"quantity"`
}

func init() {
	mel = melody.New()
	mel.Upgrader.CheckOrigin = func(r *http.Request) bool {
		//validacion dominio servidor para permitir conexiones
		return true
	}

	mel.Upgrader.ReadBufferSize = 4096
	mel.Upgrader.WriteBufferSize = 4096
	mel.Config.MaxMessageSize = 4096
	mel.Config.MessageBufferSize = 4096

	usuarios = make(map[string]string)
}

func main() {
	eco := echo.New()
	eco.Use(middleware.Logger())
	eco.Use(middleware.Recover())

	// CORS config, con metodos permitidos
	eco.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	eco.Static("/", "public")
	eco.POST("/login", login)
	eco.GET("/ws", websocket)

	err := eco.Start(":9000")

	if err != nil {
		log.Error(err)
	}

	log.Printf("Servidor iniciado correctament...")

}

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	u := &UserRequest{}
	validUsers := loadUsers()

	err := c.Bind(u)
	if err != nil {
		log.Printf("estructura no valida: %v", err)
	}

	log.Print(u)
	isAutorized := false
	for _, v := range validUsers {
		if u.Name == v.Name && u.Password == v.Password {
			isAutorized = true
			break
		}
	}

	response := make(map[string]string)
	if !isAutorized {
		response["message"] = "usuario o contraseña no validos"
		return c.JSON(http.StatusUnauthorized, response)
	}

	response["message"] = "ok"
	response["token"] = token
	return c.JSON(http.StatusOK, response)
}

func loadUsers() []UserRequest {
	usr := make([]UserRequest, 0)

	usr = append(usr, UserRequest{"jesus", "123"}, UserRequest{"zero", "123"}, UserRequest{"fulano", "123"})

	return usr
}

func websocket(c echo.Context) error {
	if c.QueryParam("token") == token {
		mel.HandleRequest(c.Response().Writer, c.Request())
		mel.HandleConnect(conectarse)
		mel.HandleDisconnect(desconectarse)
		mel.HandleMessage(mensaje)
		return nil
	}

	response := make(map[string]string)
	response["message"] = "no autorizado para acceder al websocket"
	return c.JSON(http.StatusForbidden, response)
}

func conectarse(s *melody.Session) {
	uName := getUserFromQuery(s)
	usuarios[uName] = uName
}

func getUserFromQuery(s *melody.Session) string {
	uName := s.Request.URL.Query().Get("uname")
	return uName
}

func desconectarse(s *melody.Session) {
	uName := getUserFromQuery(s)
	delete(usuarios, uName)
}

func mensaje(s *melody.Session, msg []byte) {
	var data map[string]interface{}

	err := json.Unmarshal(msg, &data)
	if err != nil {
		log.Printf("no se pudo convertion json recibido", err)
		return
	}

	tipo, ok := data["type"]
	if !ok {
		log.Printf("el mensaje recibido no tiene tipo")
	}

	switch tipo {
	case "message":
		uName := getUserFromQuery(s)
		msg := data["message"]
		dr := DataResponse{
			Name:    uName,
			Message: msg.(string),
		}
		rws := ResponseWebSocket{
			Tipo:         "message",
			DataResponse: dr,
		}
		r, err := json.Marshal(rws)
		if err != nil {
			log.Printf("no se pudo procesar el objeto de respuesta: %v", err)
		}
		mel.Broadcast(r)
	case "sale":
		product := data["product"]
		quantity := data["quantity"]
		ds := DataSale{
			Product:  product.(float64),
			Quantity: quantity.(float64),
		}
		rws := ResponseSale{
			Tipo:     "sale",
			DataSale: ds,
		}
		r, err := json.Marshal(rws)
		if err != nil {
			log.Printf("no se pudo procesar el objeto de respuesta: %v", err)
		}
		mel.Broadcast(r)
	case "ping":
		mel.BroadcastFilter([]byte(`{"type":"pong"`),
			func(sa *melody.Session) bool {
				return sa == s
			},
		)
	default:
		log.Printf("tipo no procesable")
	}
}
