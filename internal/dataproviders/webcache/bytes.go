package webcache

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"time"
)

func (c *Cache) SetBytes(ctx context.Context, u string, v []byte) {
	err := c.setBytes(u, v)
	if err != nil {
		c.logger.Error(
			"web cache: set bytes",
			slog.Any("error", err),
		)
	}

	c.metricProvider.IncWebCacheCounter("set")
}

func (c *Cache) GetBytes(ctx context.Context, u string) ([]byte, bool) {
	v, ok, err := c.getBytes(u)
	if err != nil {
		c.logger.Error(
			"web cache: get bytes",
			slog.Any("error", err),
		)
	}

	if ok {
		c.metricProvider.IncWebCacheCounter("hit")
	} else {
		c.metricProvider.IncWebCacheCounter("miss")
	}

	return v, ok
}

func (c *Cache) setBytes(u string, v []byte) (err error) {
	name := c.fileName(u)

	f, err := os.Create(path.Join(c.baseDir, name))
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			closeErr = fmt.Errorf("close file: %w", closeErr)
		}

		err = errors.Join(err, closeErr)
	}()

	_, err = f.Write(v)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (c *Cache) getBytes(u string) (_ []byte, _ bool, err error) {
	name := c.fileName(u)

	t, ok, err := c.fileInfo(name)
	if err != nil {
		return nil, false, fmt.Errorf("check file: %w", err)
	}

	limit := time.Now().Add(-c.ttl)

	if !ok || t.Before(limit) {
		return nil, false, nil
	}

	f, err := os.Open(path.Join(c.baseDir, name))
	if err != nil {
		return nil, false, fmt.Errorf("create file: %w", err)
	}

	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			closeErr = fmt.Errorf("close file: %w", closeErr)
		}

		err = errors.Join(err, closeErr)
	}()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, false, fmt.Errorf("read file: %w", err)
	}

	return data, true, nil
}
