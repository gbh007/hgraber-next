package datafs

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
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
}

func New(logger *slog.Logger, metricProvider metricProvider, path string) (*Storage, error) {
	err := createDir(path)
	if err != nil {
		return nil, err
	}

	return &Storage{
		logger:         logger,
		metricProvider: metricProvider,

		fsPath: path,
	}, nil
}

func (s *Storage) filepath(fileID uuid.UUID) string {
	return path.Join(s.fsPath, fileID.String())
}

func createDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
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
