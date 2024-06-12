package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func TimeToDB(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t.UTC(),
		Valid: !t.IsZero(),
	}
}

func UUIDToDB(u uuid.UUID) sql.NullString {
	return sql.NullString{
		String: u.String(),
		Valid:  u == uuid.Nil,
	}
}

func StringToDB(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}
