package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	*sql.DB
}

func InitMySQL() *DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/music")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}