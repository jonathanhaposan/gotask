package server

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	dbClient "github.com/jonathanhaposan/gotask/common/db"
	client "github.com/jonathanhaposan/gotask/common/redis"
)

var (
	redisPool *redis.Pool
	db        *sql.DB
)

const (
	RequestLogin       = 1
	RequestEdit        = 2
	RequestCheckCookie = 3
	imageDirectory     = "../../file/image"
)

func InitRedisConn() {
	redisPool = client.InitRedis()
	if redisPool == nil {
		log.Println("Error connecting")
	}
}

func InitDBConn() {
	conn, err := dbClient.InitDB()
	if err != nil {
		log.Println("Error init DB Conn", err)
		return
	}

	db = conn
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
}
