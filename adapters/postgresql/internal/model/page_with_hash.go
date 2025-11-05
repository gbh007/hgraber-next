package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

type PageWithHash struct {
	pages Page
	files File
}

func NewPageWithHash(pages Page, files File) PageWithHash {
	return PageWithHash{
		pages: pages,
		files: files,
	}
}

func (p PageWithHash) JoinString() string {
	return JoinPageAndFile(p.pages, p.files)
}

func (p PageWithHash) Columns() []string {
	return []string{
		p.pages.ColumnBookID(),
		p.pages.ColumnPageNumber(),
		p.pages.ColumnExt(),
		p.pages.ColumnOriginURL(),
		p.pages.ColumnDownloaded(),
		p.pages.ColumnFileID(),
		p.pages.ColumnCreateAt(),
		p.pages.ColumnLoadAt(),
		p.files.ColumnMd5Sum(),
		p.files.ColumnSha256Sum(),
		p.files.ColumnSize(),
		p.files.ColumnFSID(),
	}
}

func (PageWithHash) Scanner(p *core.PageWithHash) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			originURL sql.NullString
			fileID    sql.NullString
			loadAt    sql.NullTime
			md5Sum    sql.NullString
			sha256Sum sql.NullString
			size      sql.NullInt64
			fsID      uuid.NullUUID
		)

		err := rows.Scan(
			&p.BookID,
			&p.PageNumber,
			&p.Ext,
			&originURL,
			&p.Downloaded,
			&fileID,
			&p.CreateAt,
			&loadAt,
			&md5Sum,
			&sha256Sum,
			&size,
			&fsID,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		if originURL.Valid {
			p.OriginURL, err = url.Parse(originURL.String)
			if err != nil {
				return fmt.Errorf("convert to entity: %w", err)
			}
		}

		if fileID.Valid {
			p.FileID, err = uuid.Parse(fileID.String)
			if err != nil {
				return fmt.Errorf("convert to entity: %w", err)
			}
		}

		p.LoadAt = loadAt.Time
		p.Md5Sum = md5Sum.String
		p.Sha256Sum = sha256Sum.String
		p.Size = size.Int64
		p.FSID = fsID.UUID

		return nil
	}
}
