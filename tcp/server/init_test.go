package server

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInitRedisConn(t *testing.T) {
	InitRedisConn()
	if redisPool == nil {
		t.Errorf("Redis not init")
	}
	redisPool.Close()
}

func TestInitDBConn(t *testing.T) {
	InitDBConn()
	if db == nil {
		t.Errorf("DB not init")
	}
	db.Close()
}
