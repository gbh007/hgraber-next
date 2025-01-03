package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"hgnext/internal/entities"
)

func (d *Database) SetDeadHash(ctx context.Context, hash entities.DeadHash) error {
	builder := squirrel.Insert("dead_hashes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"md5_sum":    hash.Md5Sum,
			"sha256_sum": hash.Sha256Sum,
			"size":       hash.Size,
			"created_at": hash.CreatedAt,
		}).
		Suffix(`ON CONFLICT DO NOTHING`)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.DeadHash, error) {
	builder := squirrel.Select(
		"md5_sum",
		"sha256_sum",
		"size",
		"created_at",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("dead_hashes").
		Where(squirrel.Eq{
			"md5_sum": md5Sums,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]entities.DeadHash, 0, len(md5Sums))

	for rows.Next() {
		hash := entities.DeadHash{}

		err = rows.Scan(
			&hash.Md5Sum,
			&hash.Sha256Sum,
			&hash.Size,
			&hash.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, hash)
	}

	return result, nil
}

func (d *Database) DeleteDeadHash(ctx context.Context, hash entities.DeadHash) error {
	builder := squirrel.Delete("dead_hashes").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"md5_sum":    hash.Md5Sum,
			"sha256_sum": hash.Sha256Sum,
			"size":       hash.Size,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
