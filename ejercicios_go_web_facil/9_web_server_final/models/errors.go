package models

import "errors"

type ValidationError error

var (
	errorUsername      = ValidationError(errors.New("Username no debe estar vacío."))
	errorShortUsername = ValidationError(errors.New("Username es demasiado corto."))
	errorLargeUsername = ValidationError(errors.New("Username es demasiado largo."))

	errorEmail = ValidationError(errors.New("Formato invalido de Email."))

	errorPasswordEncryption = ValidationError(errors.New("No es posible cifrar el texto."))

	errorLogin = ValidationError(errors.New("Usuario o password incorrectos."))
)
