package importfs

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

type metricProvider interface {
	RegisterFSActionTime(action string, fsID *uuid.UUID, d time.Duration)
}

type Storage struct {
	logger         *slog.Logger
	metricProvider metricProvider

	fsPath string

	limitOnFolder int
	fsLimitMutex  sync.Mutex

	useUnsafeCloser bool
}

func New(
	logger *slog.Logger,
	metricProvider metricProvider,
	path string,
	limitOnFolder int,
	useUnsafeCloser bool,
) (*Storage, error) {
	err := createDir(path)
	if err != nil {
		return nil, err
	}

	return &Storage{
		logger:         logger,
		metricProvider: metricProvider,

		fsPath:          path,
		limitOnFolder:   limitOnFolder,
		useUnsafeCloser: useUnsafeCloser,
	}, nil
}

func createDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if info != nil && info.IsDir() {
		return nil
	}

	if info != nil && !info.IsDir() {
		return fmt.Errorf("dir path is not dir")
	}

	err = os.MkdirAll(dirPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
