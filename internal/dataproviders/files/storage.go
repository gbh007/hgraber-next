package files

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

type Storage struct {
	fsPath string

	logger *slog.Logger
}

func New(path string, logger *slog.Logger) (*Storage, error) {
	err := createDir(path)
	if err != nil {
		return nil, err
	}

	return &Storage{
		fsPath: path,
		logger: logger,
	}, nil
}

func (s *Storage) filepath(bookID uuid.UUID, bookName string) string {
	return path.Join(s.fsPath, fmt.Sprintf(
		"%s %s.zip",
		bookID.String(),
		escapeFileName(bookName),
	))
}

func escapeFileName(n string) string {
	const (
		replacer  = ""
		maxLength = 100
	)

	// TODO: заменить на strings.Replacer
	for _, e := range []string{`\`, `/`, `|`, `:`, `"`, `*`, `?`, `<`, `>`, `.`} {
		n = strings.ReplaceAll(n, e, replacer)
	}

	if len([]rune(n)) > maxLength {
		n = string([]rune(n)[:maxLength])
	}

	return n
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
