package services

import (
	"database/sql"
	"inventory/infrastructures"
	"inventory/models"
	"inventory/repositories"
	"testing"
)

func TestGetUser(t *testing.T) {
	var expected models.UserModel
	expected = models.UserModel{}
	expected.UserId = 1
	expected.UserName = "afif0808"
	expected.UserFullname = "Muhammad Afif"
	expected.UserEmail = "afif.panai@gmail.com"

	var result models.UserModel

	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	mysqlHandler := infrastructures.MySQLHandler{mysqlConn}

	loginRepository := repositories.LoginRepository{&mysqlHandler}
	userRepository := repositories.UserRepository{&mysqlHandler}
	userService := UserService{&userRepository, &loginRepository}
	token := "token"

	result, _ = userService.GetUser(token)
	if expected != result {
		t.Errorf("Error expected %v got %v", expected, result)
	} else {
		t.Logf("Success expected %v got %v", expected, result)
	}
}
