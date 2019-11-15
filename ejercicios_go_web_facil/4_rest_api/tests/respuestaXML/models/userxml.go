package models

type UserXML struct {
	ID       int    `xml:"id"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}
