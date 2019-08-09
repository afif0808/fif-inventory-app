package controllers

import (
	"inventory/interfaces"
	"net/http"
	"strings"
)

type UserAuthorizationController struct {
	interfaces.IUserAuthorizationService
	interfaces.IUserService
}

func (controller *UserAuthorizationController) AuthenticationMiddleware(handler http.Handler) http.Handler {
	// NOTE: Serve if authenticaion succeeds
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Use http header to pass token instead , because it's more flexible
		var token string
		token = r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		var tokenCookie *http.Cookie
		var tokenCookieErr error

		handler.ServeHTTP(w, r)
	})
}
