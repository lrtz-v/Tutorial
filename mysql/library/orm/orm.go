package orm

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Config for orm connection
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

// MySQL create new connection to mysql server.
func MySQL(c *Config) (db *grom.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	return
}
