package orm

import "testing"

func TestOrmConnection(t *testing.T) {
	config := &Config{
		DSN:         "root:passwd@tcp(localhost:3306)/TestDB",
		Active:      20,
		Idle:        10,
		IdleTimeout: 3600000,
	}
	db := MySQL(config)
	db.Close()
}
