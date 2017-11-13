package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //import for sql driver
)

var db *sql.DB

func init() {
	//should be update, get config from conf
	db, _ = sql.Open("mysql", "root:iao123456@tcp(10.0.75.1:3306)/adv?charset=utf8")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
}
