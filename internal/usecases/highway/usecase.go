package highway

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

type tokenizer interface {
	New(validUntil int64) string
	Validate(token string) (int64, error)
}

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
}

type UseCase struct {
	tokenizer     tokenizer
	tokenLifeTime time.Duration
	fileStorage   fileStorage
}

func New(
	tokenizer tokenizer,
	tokenLifeTime time.Duration,
	fileStorage fileStorage,
) *UseCase {
	return &UseCase{
		tokenizer:     tokenizer,
		tokenLifeTime: tokenLifeTime,
		fileStorage:   fileStorage,
	}
}

func (uc *UseCase) NewToken(ctx context.Context) (string, time.Time, error) {
	validUntil := time.Now().Add(uc.tokenLifeTime)
	token := uc.tokenizer.New(validUntil.Unix())

	return token, validUntil, nil
}

func (uc *UseCase) ValidateToken(ctx context.Context, token string) error {
	currentTime := time.Now().Unix()
	validUntil, err := uc.tokenizer.Validate(token)
	if err != nil {
		return err
	}

	if validUntil < currentTime {
		return fmt.Errorf("expired token")
	}

	return nil
}

func (uc *UseCase) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	return uc.fileStorage.Get(ctx, fileID)
}
