package controllers

import (
	"inventory/interfaces"
	"net/http"
)

type AuthenticationController struct {
	interfaces.IAuthenticationService
}

func (controller *AuthenticationController) Authenticate(handler http.Handler) http.Handler {
	// NOTE: Serve if authenticaion succeeds
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token *http.Cookie
		// NOTE: Get the token's cookie
		token, tokenErr := r.Cookie("authtoken")

		if tokenErr != nil {
			//Return 'token not found' response
			return
		}
		tokenStr := token.Value
		_, authenticationErr := controller.GetAuthentication(tokenStr)
		if authenticationErr != nil {
			//Return error response
		}
		handler.ServeHTTP(w, r)
	})
}
