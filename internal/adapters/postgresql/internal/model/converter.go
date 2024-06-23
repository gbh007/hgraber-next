package model

import (
	"database/sql"
	"net/url"
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
		Valid:  u != uuid.Nil,
	}
}

func StringToDB(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func Int32ToDB(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: i != 0,
	}
}

func URLToDB(u *url.URL) sql.NullString {
	if u == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: u.String(),
		Valid:  true,
	}
}
