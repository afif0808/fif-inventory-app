package services

import (
	"database/sql"
	"inventory/infrastructures"
	"inventory/repositories"
	"testing"
)

func TestGetUser(t *testing.T) {
	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	mysqlHandler := infrastructures.MySQLHandler{mysqlConn}
	userRepository := repositories.UserRepository{&mysqlHandler, &ErrorService{}}
	userService := UserService{&userRepository}
	_, testErr := userService.GetUserByUsernameAndPassword("afif0808", "password")
	// aE := testErr.(models.ErrorModel)

	if testErr != nil {
		t.Error(testErr.Error())
	}
}
