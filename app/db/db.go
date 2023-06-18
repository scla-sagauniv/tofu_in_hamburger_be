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
	db_host := os.Getenv("MYSQL_HOST")
	db_port := os.Getenv("MYSQL_PORT")
	db_database := os.Getenv("MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", db_user, db_password, db_host, db_port, db_database)
	// dataSourceName := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", db_host, db_user, db_password, db_port, db_database)

	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	// Db, err = sql.Open("sqlserver", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	log.Println("connected to db")
	Db.SetMaxOpenConns(25)
}
