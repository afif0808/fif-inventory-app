package infrastructures

import (
	"github.com/gomodule/redigo/redis"
)

type RedisHandler struct {
	Conn redis.Conn
}

func (r *RedisHandler) SetCache(key string, value []byte) error {
	conn := r.Conn
	_, err := conn.Do("SET", key, value)
	return err
}
func (r *RedisHandler) GetCahce(key string) ([]byte, error) {
	conn := r.Conn
	value, err := redis.Bytes(conn.Do("GET", key))
	return value, err
}
