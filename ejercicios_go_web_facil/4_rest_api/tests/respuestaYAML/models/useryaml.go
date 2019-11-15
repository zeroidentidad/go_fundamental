package models

type UserYAML struct {
	ID       int    `yaml:"id"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
