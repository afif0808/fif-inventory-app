package repositories

import (
	"inventory/iInfrastructures"
	"inventory/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	iInfrastructures.IDbHandler
}

func (repository *UserRepository) GetUserById(userId int) (user models.UserModel, err error) {
	query, err := repository.Query("SELECT * FROM user WHERE user_id = ?", userId)
	if err != nil {
		return
	}

	for query.Next() {
		scanErr := query.Scan(&user.UserId, &user.UserName, &user.UserFullname, &user.UserEmail)
		if scanErr != nil {
			err = scanErr
			return
		}
	}

	return
}
