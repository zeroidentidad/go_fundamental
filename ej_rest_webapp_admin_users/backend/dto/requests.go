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
	RoleID          uint   `json:"role_id"`
}

type RequestRole struct {
	ID          uint     `json:"role_id,omitempty"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (u RequestUser) EmptyPassword() bool {
	return u.Password == ""
}

func (r RequestUser) ValidatePassword() *errs.AppError {
	if r.EmptyPassword() {
		return errs.NewValidationError("The password must not be empty")
	}
	if len(r.Password) < 6 {
		return errs.NewValidationError("The password is too short")
	}
	if r.Password != r.PasswordConfirm {
		return errs.NewValidationError("Invalid password confirmation")
	}
	return nil
}

func (r RequestUser) EncryptPassword() (string, *errs.AppError) {
	cost := 8
	if r.EmptyPassword() {
		return "", errs.NewValidationError("Invalid empty password")
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(r.Password), cost)
	return string(bytes), nil
}
