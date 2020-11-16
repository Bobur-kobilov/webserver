package persistence

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	db *sql.DB
}

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root1234@tcp(webserver_db_1)/webserver")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
