package dto

import (
	"encoding/json"
	"time"

	"backend/errs"

	"github.com/golang-jwt/jwt"
)

const TOKEN_DURATION = time.Hour * 720 // 1 mes
const TOKEN_SECRET = "S3cRet_n0N3_1d3Nt"

type ResponseUser struct {
	ID        uint         `json:"user_id,"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Email     string       `json:"email,omitempty"`
	RoleID    uint         `json:"role_id,omitempty"`
	Role      ResponseRole `json:"role,omitempty"`
}

type ResponseRole struct {
	ID          uint                 `json:"role_id"`
	Name        string               `json:"name"`
	Permissions []ResponsePermission `json:"permissions"`
}

type ResponsePermission struct {
	ID   uint   `json:"permission_id"`
	Name string `json:"name"`
}

type ResponseProduct struct {
	ID          uint    `json:"product_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type ResponseOrder struct {
	ID         uint                `json:"id"`
	Name       string              `json:"name"`
	Email      string              `json:"email"`
	Total      float32             `json:"total"`
	UpdatedAt  string              `json:"updated_at"`
	CreatedAt  string              `json:"created_at"`
	OrderItems []ResponseOrderItem `json:"order_items"`
}

type ResponseOrderItem struct {
	ID           uint    `json:"id"`
	OrderID      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

type ResponseSales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

type ResponseUserLogin struct {
	Token string `json:"token"`
}

type UserClaims struct {
	ResponseUser
	jwt.StandardClaims
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
		"role_id":    u.RoleID,
		"exp":        time.Now().Add(TOKEN_DURATION).Unix(),
	}

	return claims
}

func FromJwtMapClaims(mapClaims jwt.MapClaims) (uc *UserClaims, err error) {
	bytes, err := json.Marshal(mapClaims)
	if err != nil {
		return uc, err
	}

	err = json.Unmarshal(bytes, &uc)
	if err != nil {
		return uc, err
	}

	return uc, nil
}
