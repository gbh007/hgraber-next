//nolint:decorder // будет исправлено позднее
package core

import (
	"crypto/md5" //nolint:gosec // не используется для криптографии
	"crypto/sha256"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          uuid.UUID
	Filename    string
	Ext         string
	Md5Sum      string
	Sha256Sum   string
	Size        int64
	FSID        uuid.UUID
	InvalidData bool
	CreateAt    time.Time
}

func (f File) Hash() FileHash {
	return FileHash{
		Md5Sum:    f.Md5Sum,
		Sha256Sum: f.Sha256Sum,
		Size:      f.Size,
	}
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
		Md5Sum:    fmt.Sprintf("%x", md5.Sum(data)), //nolint:gosec // не используется для криптографии
		Sha256Sum: fmt.Sprintf("%x", sha256.Sum256(data)),
	}, nil
}

type SizeWithCount struct {
	Count int64
	Size  int64
}

func (s SizeWithCount) Avg() int64 {
	if s.Count == 0 {
		return 0
	}

	return s.Size / s.Count
}
