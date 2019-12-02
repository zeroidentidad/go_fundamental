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

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.username = GetStringEnv("USERNAME", "remoto")
	database.password = GetStringEnv("PASSWORD", "x1234567")
	database.host = GetStringEnv("HOST", "localhost")
	database.port = GetIntEnv("PORT", 3306)
	database.database = GetStringEnv("DATABASE", "gowebrestdb")
	database.debug = GetBoolEnv("DEBUG", true)
}

func UrlDatabase() string {
	return database.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.username, this.password, this.host, this.port, this.database)
}
