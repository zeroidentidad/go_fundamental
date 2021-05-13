package dto

import "backend/errs"

type RequestUser struct {
	ID        uint   `json:"user_id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u RequestUser) EmptyPassword() bool {
	return u.Password == ""
}

func (r RequestUser) ValidatePassword() *errs.AppError {
	if r.Password == "" {
		return errs.NewValidationError("La contraseña no debe estar vacia")
	}
	if len(r.Password) < 6 {
		return errs.NewValidationError("La contraseña es demasiado corta")
	}
	return nil
}
