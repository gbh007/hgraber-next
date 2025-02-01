package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type File struct {
	ID          string         `db:"id"`
	Filename    string         `db:"filename"`
	Ext         string         `db:"ext"`
	Md5Sum      sql.NullString `db:"md5_sum"`
	Sha256Sum   sql.NullString `db:"sha256_sum"`
	Size        sql.NullInt64  `db:"size"`
	FSID        uuid.NullUUID  `db:"fs_id"`
	InvalidData bool           `db:"invalid_data"`
	CreateAt    time.Time      `db:"create_at"`
}

func (f File) ToEntity() (entities.File, error) {
	id, err := uuid.Parse(f.ID)
	if err != nil {
		return entities.File{}, err
	}

	return entities.File{
		ID:          id,
		Filename:    f.Filename,
		Ext:         f.Ext,
		Md5Sum:      f.Md5Sum.String,
		Sha256Sum:   f.Sha256Sum.String,
		Size:        f.Size.Int64,
		FSID:        f.FSID.UUID,
		InvalidData: f.InvalidData,
		CreateAt:    f.CreateAt,
	}, nil
}
