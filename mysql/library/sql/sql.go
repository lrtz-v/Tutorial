package sql

import (
	"database/sql"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Mysql init new connection to mysql server.
func Mysql() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:passwd@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	return
}
