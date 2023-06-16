package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var Db *sql.DB

func Init() {
	db_user := os.Getenv("MYSQL_USER")
	db_password := os.Getenv("MYSQL_PASSWORD")
	db_port := os.Getenv("MYSQL_PORT")
	db_database := os.Getenv("MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?charset=utf8&parseTime=true", db_user, db_password, db_port, db_database)

	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	log.Println("connected to db")
	Db.SetMaxOpenConns(25)
}
