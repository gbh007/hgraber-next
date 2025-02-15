package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type File struct {
	ID          uuid.UUID      `db:"id"`
	Filename    string         `db:"filename"`
	Ext         string         `db:"ext"`
	Md5Sum      sql.NullString `db:"md5_sum"`
	Sha256Sum   sql.NullString `db:"sha256_sum"`
	Size        sql.NullInt64  `db:"size"`
	FSID        uuid.UUID      `db:"fs_id"`
	InvalidData bool           `db:"invalid_data"`
	CreateAt    time.Time      `db:"create_at"`
}

func (f File) ToEntity() (core.File, error) {
	return core.File{
		ID:          f.ID,
		Filename:    f.Filename,
		Ext:         f.Ext,
		Md5Sum:      f.Md5Sum.String,
		Sha256Sum:   f.Sha256Sum.String,
		Size:        f.Size.Int64,
		FSID:        f.FSID,
		InvalidData: f.InvalidData,
		CreateAt:    f.CreateAt,
	}, nil
}
