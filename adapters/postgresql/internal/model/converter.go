package model

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/pkg"
)

func TimeToDB(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t.UTC(),
		Valid: !t.IsZero(),
	}
}

func UUIDToDB(u uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  u,
		Valid: u != uuid.Nil,
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

func StringsPrefix(arr []string, prefix string) []string {
	return pkg.Map(arr, func(s string) string {
		return prefix + s
	})
}
