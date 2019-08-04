package services

import (
	"fmt"
	"inventory/iModels"
	"inventory/models"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	JWT_SIGNING_METHOD *jwt.SigningMethodHMAC
	JWT_SIGNATURE_KEY  []byte
}

type JWTClaims struct {
	jwt.StandardClaims
	UserId int
}

func (service *JWTService) JWTValidate(token string) (iModels.IUserAuthorizationModel, error) {
	var userAuthorization models.UserAuthorizationModel
	var jwtValidateErr error
	var jwtToken *jwt.Token
	var jwtTokenErr error
	var claims jwt.MapClaims
	var claimsOk bool

	jwtToken, jwtTokenErr = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != service.JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return service.JWT_SIGNATURE_KEY, nil
	})

	if jwtTokenErr != nil {
		return nil, jwtTokenErr
	}

	claims, claimsOk = jwtToken.Claims.(jwt.MapClaims)

	if !claimsOk || !jwtToken.Valid {
		return nil, jwtTokenErr
	}
	// log.Println(claims)
	userAuthorization.UserId, jwtValidateErr = strconv.Atoi(fmt.Sprint(claims["UserId"]))
	return userAuthorization, jwtValidateErr
}

func (service *JWTService) JWTCreate(claims jwt.Claims) (string, error) {
	var jwtSignedToken string
	var jwtToken *jwt.Token
	jwtToken = jwt.NewWithClaims(service.JWT_SIGNING_METHOD, claims)
	jwtSignedToken, err := jwtToken.SignedString(service.JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	return jwtSignedToken, nil

}

func (service *JWTService) JWTGetSignatureKey() []byte {
	return service.JWT_SIGNATURE_KEY
}
func (service *JWTService) JWTGetSigningMethod() *jwt.SigningMethodHMAC {
	return service.JWT_SIGNING_METHOD
}
