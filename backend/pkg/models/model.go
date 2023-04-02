package models

import (
	"database/sql"
	"time"
)

type DefaultModel struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
