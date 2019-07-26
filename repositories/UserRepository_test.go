package repositories

import (
	"database/sql"
	"inventory/infrastructures"
	"testing"
)

func TestGetUserById(t *testing.T) {

	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	mysqlHandler := infrastructures.MySQLHandler{Conn: mysqlConn}
	repo := UserRepository{&mysqlHandler}
	repo.GetUserById(1)

}
