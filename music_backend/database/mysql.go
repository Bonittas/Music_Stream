package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() *sql.DB {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/your_database_name")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}