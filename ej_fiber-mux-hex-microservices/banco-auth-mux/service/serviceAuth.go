package service

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeroidentidad/fiber-hex-api-auth/domain"
	"github.com/zeroidentidad/fiber-hex-api-auth/dto"
)

type ServiceAuth interface {
	Login(dto.RequestLogin) (*string, error)
	Verify(urlParams map[string]string) (bool, error)
}

type DefaultServiceAuth struct {
	repo            domain.StorageAuth
	rolePermissions domain.RolePermissions
}

func (s DefaultServiceAuth) Login(req dto.RequestLogin) (*string, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s DefaultServiceAuth) Verify(urlParams map[string]string) (bool, error) {
	// convertir el string token en struct JWT
	if jwtToken, err := JWTokenFromString(urlParams["token"]); err != nil {
		return false, err
	} else {
		// verificar validez, tiempo de vencimiento y firma del token
		if jwtToken.Valid {
			// convertir token claims a jwt.MapClaims
			mapClaims := jwtToken.Claims.(jwt.MapClaims)
			// convertir token claims a Claims struct
			if claims, err := domain.BuildClaimsFromJwtMapClaims(mapClaims); err != nil {
				return false, err
			} else {
				// si el rol es "user", verificar si cuenta_id y el cliente_id
				// que vienen en URL pertenecen al mismo token
				if claims.IsUserRole() {
					if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
						return false, nil
					}
				}
				// verificar que el rol esté autorizado para usar la ruta
				isAuthorized := s.rolePermissions.IsAuthorizedFor(claims.Rol, urlParams["routeName"])
				return isAuthorized, nil
			}
		} else {
			return false, errors.New("Invalid token")
		}
	}
}

func JWTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewServiceAuth(repo domain.StorageAuth, permissions domain.RolePermissions) DefaultServiceAuth {
	return DefaultServiceAuth{repo, permissions}
}
