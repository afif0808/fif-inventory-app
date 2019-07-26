package repositories

import (
	"inventory/interfaces"
	"inventory/modelInterfaces"
	"inventory/models"

	_ "github.com/go-sql-driver/mysql"
)

type LoginRepository struct {
	interfaces.IDbHandler
}

func (repository *LoginRepository) GetLoginByToken(token string) (iLogin modelInterfaces.ILoginModel, err error) {
	query, err := repository.Query("SELECT user_id FROM login WHERE login_token = '" + token + "' ")
	if err != nil {
		return
	}
	var login models.LoginModel
	for query.Next() {
		query.Scan(&login.UserId)
	}
	iLogin = login
	return
}
