package agentFS

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/google/uuid"
)

type agentController interface {
	FSCreate(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID, body io.Reader) error
	FSDelete(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) error
	FSGet(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) (io.Reader, error)
	FSIDs(ctx context.Context, agentID uuid.UUID) ([]uuid.UUID, error)
}

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type Storage struct {
	agentID uuid.UUID

	logger          logger
	agentController agentController
}

func New(
	agentID uuid.UUID,
	logger logger,
	agentController agentController,
) *Storage {
	return &Storage{
		agentID:         agentID,
		logger:          logger,
		agentController: agentController,
	}
}

func (s *Storage) Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error {
	err := s.agentController.FSCreate(ctx, s.agentID, fileID, body)
	if err != nil {
		return fmt.Errorf("agent fs: %w", err)
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID) error {
	err := s.agentController.FSDelete(ctx, s.agentID, fileID)
	if err != nil {
		return fmt.Errorf("agent fs: %w", err)
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	body, err := s.agentController.FSGet(ctx, s.agentID, fileID)
	if err != nil {
		return nil, fmt.Errorf("agent fs: %w", err)
	}

	return body, nil
}

func (s *Storage) IDs(ctx context.Context) ([]uuid.UUID, error) {
	ids, err := s.agentController.FSIDs(ctx, s.agentID)
	if err != nil {
		return nil, fmt.Errorf("agent fs: %w", err)
	}

	return ids, nil
}
