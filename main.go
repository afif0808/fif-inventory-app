package main

import (
	"database/sql"
	"inventory/infrastructures"
	"inventory/repositories"
	"inventory/services"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func main() {
	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	mysqlHandler := infrastructures.MySQLHandler{Conn: mysqlConn}
	redisPool := newPool()
	redisConn := redisPool.Get()
	redisHandler := infrastructures.RedisHandler{redisConn}

	jwtService := services.JWTService{
		JWT_SIGNATURE_KEY:  []byte("ITSSIGNATUREKEY"),
		JWT_SIGNING_METHOD: jwt.SigningMethodHS256,
		ICacheHandler:      &redisHandler,
	}

	userRepository := repositories.UserRepository{&mysqlHandler}
	userService := services.UserService{&userRepository, &jwtService}

	token, err := userService.AuthenticateUser("afif0808", "password")
	cache, getErr := redis.Bool(jwtService.GetCache("BlackList" + token))
	log.Println("oi", cache, getErr)
	userAuthorizationRepository := repositories.UserAuthorizationRepository{&mysqlHandler}
	userAuthorizationService := services.UserAuthorizationService{&userAuthorizationRepository, &jwtService}

	val, err2 := userAuthorizationService.UserAuthorization(token)
	log.Println(val, err, err2)
}

//Note :
// Fix JWTService , there are messes
//
