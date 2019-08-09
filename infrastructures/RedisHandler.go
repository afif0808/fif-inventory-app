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
