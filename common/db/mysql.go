package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root@/goentry")
	if err != nil {
		log.Println("Init connection to DB :3306")
		return
	}
	return
}
