package main

import (
	"database/sql"
	"inventory/controllers"
	"inventory/infrastructures"
	"inventory/repositories"
	"inventory/services"

	"github.com/gomodule/redigo/redis"
)

type IServiceContainer interface {
	InjectUserAccessController() controllers.UserAccessController
}

type kernel struct{}

func (k *kernel) InjectUserAccessController() controllers.UserAccessController {
	var mysqlConn *sql.DB
	mysqlConn, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/inventory")
	var mysqlHandler = infrastructures.MySQLHandler{Conn: mysqlConn}

	var redisPool *redis.Pool = infrastructures.RedisNewPool()
	var redisConn redis.Conn = redisPool.Get()

	redisHandler := infrastructures.RedisHandler{Conn: redisConn}

	errorService := services.ErrorService{}

	userAccessRepository := repositories.UserAccessRepository{
		ICacheHandler: &redisHandler,
	}
	userRepository := repositories.UserRepository{
		IDbHandler: &mysqlHandler,
	}
	userService := services.UserService{
		IUserRepository: &userRepository,
	}
	jwtService := services.JWTService{
		IErrorService: &errorService,
	}
	userAccessService := services.UserAccessService{
		IUserService:            &userService,
		IJWTService:             &jwtService,
		IUserAccessRepositories: &userAccessRepository,
	}

	userAccessController := controllers.UserAccessController{&userAccessService}

	return userAccessController
}
