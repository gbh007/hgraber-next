package agentimportfs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

func (s *Storage) Get(ctx context.Context, relativePath string) (io.ReadCloser, error) {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("get_import", nil, time.Since(startAt))
	}()

	f, err := os.Open(path.Join(s.fsPath, relativePath))
	if err != nil {
		return nil, fmt.Errorf("import fs: open: %w", err)
	}

	return f, nil
}
