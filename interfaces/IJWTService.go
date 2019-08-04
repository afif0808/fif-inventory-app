package interfaces

import (
	"inventory/iModels"

	jwt "github.com/dgrijalva/jwt-go"
)

type IJWTService interface {
	JWTValidate(token string) (iModels.IUserAuthorizationModel, error)
	JWTCreate(jwt.Claims) (string, error)
	JWTGetSignatureKey() []byte
	JWTGetSigningMethod() *jwt.SigningMethodHMAC
}
