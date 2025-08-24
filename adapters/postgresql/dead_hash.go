package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (d *Database) SetDeadHash(ctx context.Context, hash core.DeadHash) error {
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

	d.SquirrelDebugLog(ctx, query, args)

	_, err = d.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) SetDeadHashes(ctx context.Context, hashes []core.DeadHash) error {
	batches := pkg.Batching(hashes, 5000)

	for _, batch := range batches {
		builder := squirrel.Insert("dead_hashes").
			PlaceholderFormat(squirrel.Dollar).
			Columns(
				"md5_sum",
				"sha256_sum",
				"size",
				"created_at",
			).
			Suffix(`ON CONFLICT DO NOTHING`)

		for _, hash := range batch {
			builder = builder.Values(hash.Md5Sum,
				hash.Sha256Sum,
				hash.Size,
				hash.CreatedAt,
			)
		}

		query, args, err := builder.ToSql()
		if err != nil {
			return fmt.Errorf("build query: %w", err)
		}

		d.SquirrelDebugLog(ctx, query, args)

		_, err = d.Pool.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("exec query: %w", err)
		}
	}

	return nil
}

func (d *Database) DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.DeadHash, error) {
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

	d.SquirrelDebugLog(ctx, query, args)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.DeadHash, 0, len(md5Sums))

	for rows.Next() {
		hash := core.DeadHash{}

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

func (d *Database) DeleteDeadHash(ctx context.Context, hash core.DeadHash) error {
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

	d.SquirrelDebugLog(ctx, query, args)

	_, err = d.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeleteDeadHashes(ctx context.Context, hashes []core.DeadHash) error {
	batch := &pgx.Batch{}

	resultCount := 0

	for _, hash := range hashes {
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

		d.SquirrelDebugLog(ctx, query, args)
		batch.Queue(query, args...)

		resultCount++
	}

	batchResult := d.Pool.SendBatch(ctx, batch)
	defer batchResult.Close()

	for range resultCount {
		_, err := batchResult.Exec()
		if err != nil {
			return fmt.Errorf("exec query: %w", err)
		}
	}

	return nil
}
