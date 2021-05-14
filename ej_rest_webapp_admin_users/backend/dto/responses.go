package dto

import (
	"time"

	"backend/errs"

	"github.com/dgrijalva/jwt-go"
)

const TOKEN_DURATION = time.Hour * 720 // 1 mes
const TOKEN_SECRET = "S3cRet_n0N3_1d3Nt"

type ResponseUser struct {
	ID        uint   `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type ResponseUserLogin struct {
	Token string `json:"token"`
}

func (u ResponseUser) CreateToken() (*ResponseUserLogin, *errs.AppError) {
	var claims jwt.MapClaims = u.claimsForUser()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(TOKEN_SECRET))
	if err != nil {
		return nil, errs.NewBadRequestError("Cannot generate token")
	}

	usrLogin := ResponseUserLogin{
		Token: signedTokenAsString,
	}

	return &usrLogin, nil
}

func (u ResponseUser) claimsForUser() jwt.MapClaims {
	claims := jwt.MapClaims{
		"user_id":    u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"email":      u.Email,
		"exp":        time.Now().Add(TOKEN_DURATION).Unix(),
	}
	return claims
}
