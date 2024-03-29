package iServices

import (
	"inventory/iModels"

	jwt "github.com/dgrijalva/jwt-go"
)

type IJWTService interface {
	JWTAuthenticate(token string) (iModels.IUserAccessModel, error)
	JWTCreate(jwt.Claims) (string, error)
	JWTGetSignatureKey() []byte
	JWTGetSigningMethod() *jwt.SigningMethodHMAC
}
