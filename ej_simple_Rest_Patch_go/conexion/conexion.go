package conexion

import (
	"log"

	"../estructuras"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

const engine string = "mysql"
const user string = "remoto"
const pass string = "x1234567"
const dbname string = "go_rest_patch"

func InitDB() {
	conn = ConnORM(ConnString())
	log.Println("Conexion correcta")
}

func ConnORM(stringConn string) *gorm.DB {
	conn, err := gorm.Open(engine, stringConn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conn
}

func ConnString() string {
	return user + ":" + pass + "@/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
}

func ConnClose() {
	conn.Close()
	log.Println("Conexion cerrada")
}

//Funcs consultas

func GetUser(id string) estructuras.Usuario {
	usuario := estructuras.Usuario{}
	conn.Where("id=?", id).First(&usuario)

	return usuario
}
