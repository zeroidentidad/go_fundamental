package config

import "fmt"

type Config interface {
	url() string
}

type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
}

type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	database.username = GetStringEnv("USERNAME", "remoto")
	database.password = GetStringEnv("PASSWORD", "x1234567")
	database.host = GetStringEnv("HOST", "localhost")
	database.port = GetIntEnv("PORT", 3306)
	database.database = GetStringEnv("DATABASE", "gowebrestdb")
	database.debug = GetBoolEnv("DEBUG", true)

	server = &ServerConfig{}
	server.host = GetStringEnv("HOST", "localhost")
	server.port = GetIntEnv("PORT", 8080)
	server.debug = GetBoolEnv("DEBUG", true)
}

func UrlDatabase() string {
	return database.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.username, this.password, this.host, this.port, this.database)
}

func UrlServer() string {
	return server.url()
}

func ServerPort() int {
	return server.port
}

func (this *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}

/* funcs entorno de desarrollo y pruebas */

func Debug() bool {
	//return database.debug
	return server.debug
}

/* funcs dir templates src html */

func DirTemplate() string {
	return "templates/**/*.html"
}

func DirTemplateError() string {
	return "templates/error.html"
}
