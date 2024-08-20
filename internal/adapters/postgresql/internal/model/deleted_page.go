package model

import (
	"database/sql"
	"time"
)

type DeletedPage struct {
	BookID     string         `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	OriginURL  sql.NullString `db:"origin_url"`
	Md5Sum     sql.NullString `db:"md5_sum"`
	Sha256Sum  sql.NullString `db:"sha256_sum"`
	Size       sql.NullInt64  `db:"size"`
	Downloaded bool           `db:"downloaded"`
	CreatedAt  time.Time      `db:"created_at"`
	LoadedAt   sql.NullTime   `db:"loaded_at"`
	DeletedAt  sql.NullTime   `db:"deleted_at"`
}
