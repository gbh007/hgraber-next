package fileStorage

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/adapters/agentFS"
	"hgnext/internal/adapters/localFiles"
	"hgnext/internal/entities"
)

type agentController interface {
	FSCreate(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID, body io.Reader) error
	FSDelete(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) error
	FSGet(ctx context.Context, agentID uuid.UUID, fileID uuid.UUID) (io.Reader, error)
	FSIDs(ctx context.Context, agentID uuid.UUID) ([]uuid.UUID, error)
}

type dataStorage interface {
	File(ctx context.Context, id uuid.UUID) (entities.File, error)
}

type legacyFileStorage interface {
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
	Delete(ctx context.Context, fileID uuid.UUID) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	IDs(ctx context.Context) ([]uuid.UUID, error)
}

type Storage struct {
	logger          *slog.Logger
	agentController agentController
	dataStorage     dataStorage

	legacyFileStorage legacyFileStorage
}

func New(
	logger *slog.Logger,
	agentController agentController,
	dataStorage dataStorage,
) *Storage {
	return &Storage{
		logger:          logger,
		agentController: agentController,
		dataStorage:     dataStorage,
	}
}

func (s *Storage) InitLegacy(ctx context.Context, fsAgentID uuid.UUID, filePath string) error {
	var err error

	switch {
	case fsAgentID != uuid.Nil:
		s.legacyFileStorage = agentFS.New(fsAgentID, s.logger, s.agentController)

		s.logger.DebugContext(
			ctx, "use agent file storage",
			slog.String("agent_id", fsAgentID.String()),
		)

	case filePath != "":
		s.legacyFileStorage, err = localFiles.New(filePath, s.logger)
		if err != nil {
			return fmt.Errorf("fail init local file storage: %w", err)
		}

		s.logger.DebugContext(
			ctx, "use local file storage",
			slog.String("path", filePath),
		)

	default:
		return fmt.Errorf("no configuration for file storage")
	}

	return nil
}

func (s *Storage) FSIDForDownload(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil // FIXME: реализовать поход в БД.
}

func (s *Storage) Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error {
	return s.legacyFileStorage.Create(ctx, fileID, body)
}

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) error {
	return s.legacyFileStorage.Delete(ctx, fileID)
}

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error) {
	return s.legacyFileStorage.Get(ctx, fileID)
}

func (s *Storage) IDs(ctx context.Context, fsID uuid.UUID) ([]uuid.UUID, error) {
	return s.legacyFileStorage.IDs(ctx)
}
