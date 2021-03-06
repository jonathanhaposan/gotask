package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root@tcp(mysql.db:3306)/goentry")
	if err != nil {
		log.Printf("[db][InitDB]Failed to open connection. %+v\n", err)
		return
	}
	return
}
