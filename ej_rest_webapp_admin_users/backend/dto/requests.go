package dto

import (
	"backend/errs"

	"golang.org/x/crypto/bcrypt"
)

type RequestUser struct {
	ID              uint   `json:"user_id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty"`
}

func (u RequestUser) EmptyPassword() bool {
	return u.Password == ""
}

func (r RequestUser) ValidatePassword() *errs.AppError {
	if r.EmptyPassword() {
		return errs.NewValidationError("La contraseña no debe estar vacia")
	}
	if len(r.Password) < 6 {
		return errs.NewValidationError("La contraseña es demasiado corta")
	}
	if r.Password != r.PasswordConfirm {
		return errs.NewValidationError("Confirmación de contraseña no válida")
	}
	return nil
}

func (r RequestUser) EncryptPassword() (string, *errs.AppError) {
	cost := 8
	if r.EmptyPassword() {
		return "", errs.NewValidationError("contraseña vacia no puede encriptase")
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(r.Password), cost)
	return string(bytes), nil
}
