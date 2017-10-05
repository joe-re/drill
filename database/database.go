package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// DbExec execute any SQL
func DbExec(db *sql.DB, q string, args ...interface{}) sql.Result {
	var result, err = db.Exec(q, args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return result
}

// Query execute SQL select
func Query(db *sql.DB, q string, args ...interface{}) *sql.Rows {
	var rows, err = db.Query(q, args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rows
}

func Connect() *sql.DB {
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}
