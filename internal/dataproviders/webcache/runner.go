package webcache

import (
	"context"
	"log/slog"
	"os"
	"path"
	"strings"
	"time"
)

func (c *Cache) Name() string {
	return "web cache"
}

func (c *Cache) Start(parentCtx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		c.logger.InfoContext(parentCtx, "web cache start")
		defer c.logger.InfoContext(parentCtx, "web cache stop")

		t := time.NewTicker(c.cleanInterval)
		defer t.Stop()

		for {
			select {
			case <-parentCtx.Done():
				c.logger.InfoContext(parentCtx, "stopping web cache")

				return

			case <-t.C:
				fileList, err := os.ReadDir(c.baseDir)
				if err != nil {
					c.logger.Error(
						"web cache clean: read dir",
						slog.Any("error", err),
					)

					continue
				}

				limit := time.Now().Add(-c.ttl)

				for _, fileInfo := range fileList {
					name := fileInfo.Name()

					if fileInfo.IsDir() || !strings.HasSuffix(name, fileSuffix) {
						continue
					}

					t, ok, err := c.fileInfo(name)
					if err != nil {
						c.logger.Error(
							"web cache clean: get file info",
							slog.Any("error", err),
						)

						continue
					}

					if !ok || t.After(limit) {
						continue
					}

					err = os.Remove(path.Join(c.baseDir, name))
					if err != nil {
						c.logger.Error(
							"web cache clean: remove file",
							slog.Any("error", err),
						)

						continue
					}

					c.metricProvider.IncWebCacheCounter("expire")
				}

			}
		}
	}()

	return done, nil
}
