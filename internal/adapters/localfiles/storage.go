package localfiles

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/google/uuid"
)

type Storage struct {
	fsPath string

	logger *slog.Logger
}

func New(dirPath string, logger *slog.Logger) (*Storage, error) {
	err := createDir(dirPath)
	if err != nil {
		return nil, err
	}

	return &Storage{
		fsPath: dirPath,
		logger: logger,
	}, nil
}

func (s *Storage) filepath(fileID uuid.UUID) string {
	return path.Join(s.fsPath, fileID.String())
}

func createDir(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("os stat: %w", err)
	}

	if info != nil && !info.IsDir() {
		return errors.New("dir path is not dir")
	}

	err = os.MkdirAll(dirPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	return nil
}
