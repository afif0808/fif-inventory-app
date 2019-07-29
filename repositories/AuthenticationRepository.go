package repositories

import (
	"inventory/iInfrastructures"
	"inventory/iModels"
	"inventory/models"
)

type AuthenticationRepository struct {
	iInfrastructures.IDbHandler
}

// NOTE: Function to get only a row
func (repository *AuthenticationRepository) GetAuthenticationByToken(token string) (iAuthentication iModels.IAuthenticationModel, err error) {
	query := repository.QuerySingle("SELECT user_id FROM authentication WHERE authentication_token = '" + token + "' ")

	if err != nil {
		return models.AuthenticationModel{}, err
	}
	// var login models.LoginModel
	var authentication models.AuthenticationModel

	var scanErr error
	scanErr = query.Scan(&authentication.UserId)
	if scanErr != nil {
		return models.AuthenticationModel{}, scanErr
	}
	iAuthentication = authentication
	return
}
