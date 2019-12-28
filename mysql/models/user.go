package models

import (
	"time"

	"github.com/jinzhu/gorm"

	"database/sql"
)

type User struct {
	gorm.Model
	ID           int64 `gorm:"primary_key"`
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	CreatedAt    time.Time                        // column name is `created_at`
}
