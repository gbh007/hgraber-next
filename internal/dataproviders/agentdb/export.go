package agentdb

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/google/uuid"
)

func (s *Storage) CreateImport(ctx context.Context, info agentdto.ImportInfo) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO import_infos (book_id, book_url, relative_path, imported_at) VALUES (?,?,?,?) ON CONFLICT DO NOTHING;`,
		info.BookID,
		URLToDB(info.BookURL),
		info.FSPath,
		info.ImportedAt.Unix(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ImportedCountByID(ctx context.Context, bookID uuid.UUID) (int, error) {
	var c sql.NullInt64

	err := s.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM import_infos WHERE book_id = ?;`, bookID)
	if err != nil {
		return 0, err
	}

	return int(c.Int64), nil
}

func (s *Storage) ImportedCountByURL(ctx context.Context, u url.URL) (int, error) {
	var c sql.NullInt64

	err := s.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM import_infos WHERE book_url = ?;`, u.String())
	if err != nil {
		return 0, err
	}

	return int(c.Int64), nil
}

func (s *Storage) ImportedCountByRelativePath(ctx context.Context, path string) (int, error) {
	var c sql.NullInt64

	err := s.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM import_infos WHERE relative_path = ?;`, path)
	if err != nil {
		return 0, err
	}

	return int(c.Int64), nil
}
