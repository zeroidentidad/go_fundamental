package utils

import (
	"regexp"

	"postgres-jwt-react/db"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

// ValidateUser devuelve slice de cadena de errores de validación
func ValidateUser(user db.Register, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck != true {
		err = append(err, "Email inválido")
	}
	if len(user.Password) < 4 {
		err = append(err, "Contraseña no válida, la contraseña debe tener más de 4 caracteres")
	}
	if len(user.Name) < 1 {
		err = append(err, "Nombre inválido, ingrese un nombre")
	}

	return err
}

func ValidatePasswordReset(resetPassword db.ResetPassword) (bool, string) {
	if len(resetPassword.Password) < 4 {
		return false, "Contraseña no válida, la contraseña debe tener más de 4 caracteres"
	}
	if resetPassword.Password != resetPassword.ConfirmPassword {
		return false, "Error al restablecer contraseña, las contraseñas deben coincidir "
	}
	return true, ""
}
