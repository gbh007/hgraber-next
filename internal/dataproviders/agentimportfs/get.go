package agentimportfs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (s *Storage) Get(ctx context.Context, relativePath string) (io.Reader, error) {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("get_import", nil, time.Since(startAt))
	}()

	f, err := os.Open(path.Join(s.fsPath, relativePath))
	if err != nil {
		return nil, fmt.Errorf("import fs: open: %w", err)
	}

	if s.useUnsafeCloser {
		return pkg.UnsafeCloser(f), nil
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("import fs: read all: %w", err)
	}

	return bytes.NewReader(data), nil
}
