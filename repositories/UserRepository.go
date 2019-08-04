package repositories

import (
	"database/sql"
	"inventory/iInfrastructures"
	"inventory/iModels"
	"inventory/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	iInfrastructures.IDbHandler
}

func (repository *UserRepository) GetUserById(userId int) (iModels.IUserModel, error) {
	var user models.UserModel
	query := repository.QuerySingle("SELECT * FROM user WHERE user_id = ?", userId)
	scanErr := query.Scan(&user.UserId, &user.UserName, &user.UserFullname, &user.UserEmail)
	if scanErr != nil && scanErr != sql.ErrNoRows {
		return nil, scanErr
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
	if scanErr != nil && scanErr != sql.ErrNoRows {
		return nil, scanErr
	}
	return &user, nil
}
