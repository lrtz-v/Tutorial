package sql

import (
	"testing"
)

func TestConnect(t *testing.T) {
	config := &Config{
		DSN:         "root:passwd@tcp(localhost:3306)/TestDB",
		Active:      20,
		Idle:        10,
		IdleTimeout: 3600000,
	}

	db := Mysql(config)
	if err := db.Ping(); err != nil {
		t.Fatalf("Ping err: %s", err)
	}
	db.Close()
}
