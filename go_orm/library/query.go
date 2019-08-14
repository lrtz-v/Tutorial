package library

import (
	"database/sql"

	common "github.com/Tutorials/go_orm/common"
	"github.com/andrewpillar/query"
	_ "github.com/andrewpillar/query"
)

// Query .
type Query struct {
	db *sql.DB
}

func InitDB() (*Query, error) {
	db, err := common.Connect()
	if err != nil {
		return nil, err
	}
	return &Query{db: db}, nil
}

func (Q *Query) Close() {
	Q.db.Close()
}

func (Q *Query) Query(query string) (*sql.Rows, error) {
	return Q.db.Query(query)
}

func (Q *Query) QueryV2() (*sql.Rows, error) {
	posts := make([]*Post, 0)
	q := query.Select(
		query.Columns("data_type, data_status"),
		query.Table("offline_data_sync_status"),
		query.WhereEq("id", 1417),
	)
}
