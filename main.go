package main

import (
	"database/sql"
	"inventory/infrastructures"
	"inventory/repositories"
	"inventory/services"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	mysqlHandler := infrastructures.MySQLHandler{Conn: mysqlConn}

	jwtService := services.JWTService{
		JWT_SIGNATURE_KEY:  []byte("ITSSIGNATUREKEY"),
		JWT_SIGNING_METHOD: jwt.SigningMethodHS256,
	}

	userRepository := repositories.UserRepository{&mysqlHandler}
	userService := services.UserService{&userRepository, &jwtService}

	token, err := userService.AuthenticateUser("afif0808", "password")

	userAuthorizationRepository := repositories.UserAuthorizationRepository{&mysqlHandler}
	userAuthorizationService := services.UserAuthorizationService{&userAuthorizationRepository, &jwtService}

	val, err2 := userAuthorizationService.UserAuthorization(token)

}
