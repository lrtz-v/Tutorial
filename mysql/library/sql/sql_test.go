package sql

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := Mysql()
	if err := db.Ping(); err != nil {
		t.Fatalf("Ping err: %s", err)
	}
	db.Close()
}
