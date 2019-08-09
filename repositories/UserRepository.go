package repositories

import (
	"database/sql"
	"inventory/appErr"
	"inventory/iInfrastructures"
	"inventory/iModels"
	"inventory/iServices"
	"inventory/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	iInfrastructures.IDbHandler
	iServices.ErrorService
}

func (repository *UserRepository) GetUserByIdentity(userIdentity string) (iModels.IUserModel, error) {
	var user models.UserModel
	query := repository.QuerySingle("SELECT * FROM user WHERE user_id = ?", userIdentity)
	scanErr := query.Scan(&user.UserId, &user.UserName, &user.UserFullname, &user.UserEmail)
	switch true {
	case scanErr == sql.ErrNoRows:
		return nil, repository.ErrorAndCode(scanErr, appErr.ErrNotFound)
	case scanErr != nil:
		return nil, repository.ErrorAndCode(scanErr, appErr.ErrUnexpected)
	}

	return &user, nil
}
func (repository *UserRepository) GetUserByUsernameAndPassword(username, password string) (iModels.IUserModel, error) {
	var user models.UserModel
	query := repository.QuerySingle(
		"SELECT user.* FROM user INNER JOIN userPassword ON user.user_id = userPassword.user_id WHERE user_name = ? AND userPassword_hash = ? ",
		username, password,
	)

	scanErr := query.Scan(&user.UserId, &user.UserName, &user.UserFullname, &user.UserEmail)
	// NOTE: 'empty result' is one of the error

	switch true {
	case scanErr == sql.ErrNoRows:
		return nil, repository.ErrorAndCode(scanErr, appErr.ErrNotFound)
	case scanErr != nil:
		return nil, repository.ErrorAndCode(scanErr, appErr.ErrUnexpected)
	}

	return &user, nil
}
