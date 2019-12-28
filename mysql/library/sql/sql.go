package sql

import (
	"database/sql"
	"time"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Config for orm connection
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

// Mysql init new connection to mysql server.
func Mysql(c *Config) (db *sql.DB) {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	db.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return
}
