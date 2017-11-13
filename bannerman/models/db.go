package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //mysql driver for db connect
)

var db *sql.DB

func init() {
	dbSourceName := "root:iao123456@tcp(10.0.75.1:3306)/adv?charset=utf8"
	db, _ = sql.Open("mysql", dbSourceName)
	db.SetMaxOpenConns(10)
	db.SetMaxOpenConns(3)
	//db.Ping() //trigger db connect
}
