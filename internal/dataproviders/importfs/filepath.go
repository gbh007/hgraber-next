package importfs

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (s *Storage) filepath(bookID uuid.UUID, bookName string) (relativePath, absolutePath string) {
	relativePath = s.relativePath(bookID, bookName)
	absolutePath = path.Join(s.fsPath, relativePath)

	return
}

func (s *Storage) relativePath(bookID uuid.UUID, bookName string) string {
	return bookID.String() + " " + escapeFileName(bookName) + ".zip"
}

func (s *Storage) filepathWithLimits(bookID uuid.UUID, bookName string) (relativePath, absolutePath string, err error) {
	relativePath = s.relativePath(bookID, bookName)

	prefix, err := s.limitPathPrefix()
	if err != nil {
		return "", "", err
	}

	relativePath = path.Join(prefix, relativePath)
	absolutePath = path.Join(s.fsPath, relativePath)

	return relativePath, absolutePath, nil
}

func (s *Storage) limitPathPrefix() (string, error) {
	basePrefix := time.Now().Format("2006-01-02") + "-p"

	// TODO: вынести ограничение итераций в конфигурацию
	for i := 0; i < 10000; i++ {
		prefix := basePrefix + strconv.Itoa(i)

		c, err := s.limitFromFolder(path.Join(s.fsPath, prefix))
		if err != nil {
			return "", fmt.Errorf("check limit from folder: %w", err)
		}

		if c < s.limitOnFolder {
			return prefix, nil
		}
	}

	return "", fmt.Errorf("empty folder not found")
}

func (s *Storage) limitFromFolder(dirPath string) (int, error) {
	info, err := os.Stat(dirPath)
	if errors.Is(err, os.ErrNotExist) {
		err := createDir(dirPath)
		if err != nil {
			return 0, fmt.Errorf("create new dir: %w", err)
		}

		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf("check dir: %w", err)
	}

	if !info.IsDir() {
		return 0, fmt.Errorf("not dir in %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return 0, fmt.Errorf("read dir: %w", err)
	}

	var c int

	for _, entry := range entries {
		if !entry.IsDir() && path.Ext(entry.Name()) == ".zip" {
			c++
		}
	}

	return c, nil
}

func escapeFileName(n string) string {
	const (
		replacer  = ""
		maxLength = 100
	)

	// TODO: заменить на strings.Replacer
	for _, e := range []string{`\`, `/`, `|`, `:`, `"`, `*`, `?`, `<`, `>`, `.`, "\t"} {
		n = strings.ReplaceAll(n, e, replacer)
	}

	if len([]rune(n)) > maxLength {
		n = string([]rune(n)[:maxLength])
	}

	return n
}
