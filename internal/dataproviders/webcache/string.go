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

func (c *Cache) SetString(ctx context.Context, u, v string) {
	err := c.setString(u, v)
	if err != nil {
		c.logger.Error(
			"web cache: set string",
			slog.Any("error", err),
		)
	}

	c.metricProvider.IncWebCacheCounter("set")
}

func (c *Cache) GetString(ctx context.Context, u string) (string, bool) {
	v, ok, err := c.getString(u)
	if err != nil {
		c.logger.Error(
			"web cache: get string",
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

func (c *Cache) setString(u, v string) (err error) {
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

	_, err = f.WriteString(v)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (c *Cache) getString(u string) (_ string, _ bool, err error) {
	name := c.fileName(u)

	t, ok, err := c.fileInfo(name)
	if err != nil {
		return "", false, fmt.Errorf("check file: %w", err)
	}

	limit := time.Now().Add(-c.ttl)

	if !ok || t.Before(limit) {
		return "", false, nil
	}

	f, err := os.Open(path.Join(c.baseDir, name))
	if err != nil {
		return "", false, fmt.Errorf("create file: %w", err)
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
		return "", false, fmt.Errorf("read file: %w", err)
	}

	return string(data), true, nil
}
