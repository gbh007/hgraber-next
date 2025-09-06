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
		Int32: int32(i), //nolint:gosec // для удобства, гарантию обеспечивает логика сверху
		Valid: i != 0,
	}
}

func Int64ToDB(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
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

// TODO: возможно стоит перенести в pkg
func StringsPrefix(arr []string, prefix string) []string {
	return pkg.Map(arr, func(s string) string {
		return prefix + s
	})
}

func NilInt64ToDB(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: *i,
		Valid: true,
	}
}

func NilInt64FromDB(i sql.NullInt64) *int64 {
	if !i.Valid {
		return nil
	}

	return &i.Int64
}
