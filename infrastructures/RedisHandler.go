package infrastructures

import (
	"github.com/gomodule/redigo/redis"
)

type RedisHandler struct {
	Conn redis.Conn
}

func (r *RedisHandler) SetCache(key string, value interface{}) error {
	conn := r.Conn
	_, err := conn.Do("SET", key, value)
	return err
}
func (r *RedisHandler) DeleteCache(key string) error {
	conn := r.Conn
	_, err := conn.Do("DEL", key)
	return err
}

func (r *RedisHandler) GetCache(key string) (interface{}, error) {
	conn := r.Conn
	value, err := conn.Do("GET", key)
	return value, err

}

func RedisNewPool() *redis.Pool {
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
