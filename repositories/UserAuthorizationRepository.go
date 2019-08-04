package repositories

import (
	"database/sql"
	"inventory/iInfrastructures"
	"inventory/iModels"
	"inventory/models"
)

type UserAuthorizationRepository struct {
	iInfrastructures.IDbHandler
}

// NOTE: Function to get only a row
func (repository *UserAuthorizationRepository) GetUserAuthorizationByTokenFromDB(token string) (iModels.IUserAuthorizationModel, error) {
	query := repository.QuerySingle("SELECT user_id FROM userAuthorization WHERE userAuthorization_token = ? ", token)
	// var login models.LoginModel
	var userAuthorization models.UserAuthorizationModel

	var scanErr error
	scanErr = query.Scan(&userAuthorization.UserId)
	if scanErr != nil && scanErr != sql.ErrNoRows {
		return nil, scanErr
	}

	return &userAuthorization, nil
}
