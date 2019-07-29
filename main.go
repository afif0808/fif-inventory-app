package main

import (
	"database/sql"
	"inventory/infrastructures"
	"inventory/repositories"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	mysqlHandler := infrastructures.MySQLHandler{Conn: mysqlConn}
	repo := repositories.UserRepository{&mysqlHandler}
	_, err := repo.GetUserById(1)
	log.Println(err)
}
