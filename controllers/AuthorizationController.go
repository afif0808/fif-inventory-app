package controllers

import (
	"inventory/iModels"
	"inventory/interfaces"
	"net/http"
)

type UserAuthorizationController struct {
	interfaces.IUserAuthorizationService
	interfaces.IUserService
}

func (controller *UserAuthorizationController) AuthenticationMiddleware(handler http.Handler) http.Handler {
	// NOTE: Serve if authenticaion succeeds
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Use http header to pass token instead , because it's more flexible

		var tokenCookie *http.Cookie
		var tokenCookieErr error
		var token string

		var userAuthorizationUserId int
		var userAuthorizationUser iModels.IUserModel
		var userAuthorizationUserErr error

		tokenCookie, tokenCookieErr = r.Cookie("authtoken")

		// NOTE: Get the token's cookie

		if tokenCookieErr != nil {
			// w.Header().Set("Code", value)
			//Return 'token not found' response
			return
		}
		token = tokenCookie.Value
		userAuthorization, userAuthorizationErr := controller.
		switch true {
		case userAuthorizationErr != nil:
			// Error response
			return
		case userAuthorization != nil:
			//Authentication not found error
			return
		}
		userAuthorizationUserId = userAuthorization.GetUserAuthorizationUserId()
		userAuthorizationUser, userAuthorizationUserErr = controller.GetUser(userAuthorizationUserId)
		switch true {
		case userAuthorizationUserErr != nil:
			return
		case userAuthorizationUser != nil:
			return
		}
		handler.ServeHTTP(w, r)
	})
}
