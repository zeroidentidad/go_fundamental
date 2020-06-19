package models

/*LoginResponse tiene token que se devuelve de login */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
