package entities

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID
	Filename  string
	Ext       string
	Md5Sum    string
	Sha256Sum string
	Size      int64
	CreateAt  time.Time
}

type FileHash struct {
	Md5Sum    string
	Sha256Sum string
	Size      int64
}

func HashFile(body io.Reader) (FileHash, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return FileHash{}, fmt.Errorf("read body for hashing: %w", err)
	}

	return FileHash{
		Size:      int64(len(data)),
		Md5Sum:    fmt.Sprintf("%x", md5.Sum(data)),
		Sha256Sum: fmt.Sprintf("%x", sha256.Sum256(data)),
	}, nil
}
