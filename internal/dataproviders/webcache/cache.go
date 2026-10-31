package webcache

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strconv"
	"time"
)

const fileSuffix = ".ac"

type metricProvider interface {
	IncWebCacheCounter(action string)
}

type Cache struct {
	baseDir        string
	logger         *slog.Logger
	metricProvider metricProvider
	ttl            time.Duration
	cleanInterval  time.Duration
}

func New(
	baseDir string,
	logger *slog.Logger,
	ttl time.Duration,
	cleanInterval time.Duration,
	metricProvider metricProvider,
) (*Cache, error) {
	if baseDir == "" {
		baseDir = path.Join(os.TempDir(), "hgraber-next-agent")
	}

	err := createDir(baseDir)
	if err != nil {
		return nil, err
	}

	return &Cache{
		baseDir:        baseDir,
		logger:         logger,
		ttl:            ttl,
		cleanInterval:  cleanInterval,
		metricProvider: metricProvider,
	}, nil
}

func (c *Cache) fileName(u string) string {
	buff := bytes.Buffer{}

	md5Sum := md5.Sum([]byte(u))
	sha256Sum := sha256.Sum256([]byte(u))

	buff.WriteString(hex.EncodeToString(md5Sum[:]))
	buff.WriteString("_")
	buff.WriteString(hex.EncodeToString(sha256Sum[:]))
	buff.WriteString("_")
	buff.WriteString(strconv.Itoa(len(u)))
	buff.WriteString(fileSuffix)

	return buff.String()
}

func (c *Cache) fileInfo(name string) (time.Time, bool, error) {
	info, err := os.Stat(path.Join(c.baseDir, name))
	if errors.Is(err, os.ErrNotExist) {
		return time.Time{}, true, nil
	}

	if err != nil {
		return time.Time{}, false, err
	}

	return info.ModTime(), true, nil
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
