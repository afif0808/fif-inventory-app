package services

import (
	"fmt"
	"inventory/appErr"
	"inventory/iModels"
	"inventory/iServices"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	JWT_SIGNING_METHOD *jwt.SigningMethodHMAC
	JWT_SIGNATURE_KEY  []byte
	// iInfrastructures.ICacheHandler
	iServices.IErrorService
}

type JWTClaims struct {
	jwt.StandardClaims
	UserIdentity string
}

func (service *JWTService) JWTAuthenticate(token string) (iModels.IUserAccessModel, error) {
	// Handling error is suck
	var userAccess iModels.IUserAccessModel
	var jwtAuthenticateErr error
	var jwtToken *jwt.Token
	// var jwtTokenSignedString string
	var claims jwt.MapClaims
	var claimsOk bool

	jwtToken, jwtAuthenticateErr = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, service.ErrorAndCode(fmt.Errorf("Signing method invalid"), appErr.ErrUnexpected)
		} else if method != service.JWT_SIGNING_METHOD {
			return nil, service.ErrorAndCode(fmt.Errorf("Signing method invalid"), appErr.ErrUnexpected)
		}

		return service.JWT_SIGNATURE_KEY, nil
	})

	if jwtAuthenticateErr != nil {
		return nil, service.ErrorAndCode(jwtAuthenticateErr, appErr.ErrUnexpected)
	}

	claims, claimsOk = jwtToken.Claims.(jwt.MapClaims)
	if !claimsOk {
		return nil, service.ErrorAndCode(fmt.Errorf("Error : the jwt claims is not 'jwt.MapClaims' type"), appErr.ErrUnexpected)
	}
	// jwtTokenSignedString, jwtAuthenticateErr = jwtToken.SignedString(service.JWT_SIGNATURE_KEY)

	userAccess.SetUserIdentity(claims["UserId"])
	return userAccess, nil
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
