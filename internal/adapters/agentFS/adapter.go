package agentFS

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type agentController interface {
	FSCreate(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID, body io.Reader) error
	FSDelete(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) error
	FSGet(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) (io.Reader, error)
	FSState(ctx context.Context, agentID uuid.UUID, includeFileIDs, includeFileSizes bool) (entities.FSState, error)
}

type Storage struct {
	agentID uuid.UUID

	logger          *slog.Logger
	agentController agentController
}

func New(
	agentID uuid.UUID,
	logger *slog.Logger,
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

func (s *Storage) State(ctx context.Context, includeFileIDs, includeFileSizes bool) (entities.FSState, error) {
	state, err := s.agentController.FSState(ctx, s.agentID, includeFileIDs, includeFileSizes)
	if err != nil {
		return entities.FSState{}, fmt.Errorf("agent fs: %w", err)
	}

	return state, nil
}
