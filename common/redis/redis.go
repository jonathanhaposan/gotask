package redis

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	RedisHost = ":6379"
)

func InitRedis() (Pool *redis.Pool) {
	Pool = newPool(RedisHost)
	return
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				log.Println("Redis Error", err)
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Println("Error Ping Redis", err)
				return err
			}
			return err
		},
	}
}
