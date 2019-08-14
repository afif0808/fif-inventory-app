package controllers

import (
	"inventory/appErr"
	"inventory/iModels"
	"inventory/iServices"
	"net/http"
	"strings"
)

type UserAccessController struct {
	// iServices.IUserService
	iServices.IUserAccessService
}

func (controller *UserAccessController) UserAccessValidationMiddleware(handler http.Handler) http.Handler {
	// NOTE: Serve if authenticaion succeeds/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Use http header to pass token instead , because it's more flexible
		// var userAccess iModels.IUserAccessModel
		var validateUserErr error
		authorizationBearer := r.Header.Get("Authorization")
		token := strings.Replace(authorizationBearer, "Bearer", "", -1)
		_, validateUserErr = controller.ValidateUser(token)
		switch validateUserErr.(iModels.IErrorModel).ErrorCode() {
		case appErr.ErrNotFound:
			http.Error(w, "The requested source not found", appErr.ErrNotFound)
			return
		case appErr.ErrUnexpected:
			http.Error(w, "The requested source not found", appErr.ErrNotFound)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
