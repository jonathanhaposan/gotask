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
	imageDirectory     = "file/image"
	imageURL           = "/image/"
)

func InitRedisConn() {
	redisPool = client.InitRedis()
	if redisPool == nil {
		log.Printf("[server][InitRedisConn]Failed to get redis pool instance")
		return
	}
}

func InitDBConn() {
	log.Println("Start connecting to MySQL on :3306")
	conn, err := dbClient.InitDB()
	if err != nil {
		log.Printf("[server][InitDBConn]Failed when connecting to database. %+v\n", err)
		return
	}

	if conn == nil {
		log.Printf("[server][InitDBConn]Failed to get database connection")
		return
	}

	db = conn
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)

}
