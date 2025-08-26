package filestorage

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

type agentController interface {
	FSCreate(ctx context.Context, agentID, fileID uuid.UUID, body io.Reader) error
	FSDelete(ctx context.Context, agentID, fileID uuid.UUID) error
	FSGet(ctx context.Context, agentID, fileID uuid.UUID) (io.Reader, error)
	FSState(ctx context.Context, agentID uuid.UUID, includeFileIDs, includeFileSizes bool) (fsmodel.FSState, error)

	CreateHighwayToken(ctx context.Context, agentID uuid.UUID) (string, time.Time, error)
}

type dataStorage interface {
	File(ctx context.Context, id uuid.UUID) (core.File, error)
	FileStorages(ctx context.Context) ([]fsmodel.FileStorageSystem, error)
	FileStorage(ctx context.Context, id uuid.UUID) (fsmodel.FileStorageSystem, error)
}

type rawFileStorage interface {
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
	Delete(ctx context.Context, fileID uuid.UUID) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	State(ctx context.Context, includeFileIDs, includeFileSizes bool) (fsmodel.FSState, error)
}

type metricProvider interface {
	RegisterFSActionTime(action string, fsID *uuid.UUID, d time.Duration)
}

type rawFileStorageData struct {
	FS      rawFileStorage
	AgentID uuid.UUID
	Path    string

	EnableHighway          bool
	HighwayToken           string
	HighwayTokenValidUntil time.Time

	HighwayServerScheme       string
	HighwayServerHostWithPort string
}

type Storage struct {
	logger          *slog.Logger
	agentController agentController
	dataStorage     dataStorage
	metricProvider  metricProvider

	tryReconnect bool

	storageMap      map[uuid.UUID]rawFileStorageData
	storageMapMutex *sync.RWMutex
}

func New(
	logger *slog.Logger,
	agentController agentController,
	dataStorage dataStorage,
	metricProvider metricProvider,
	tryReconnect bool,
) *Storage {
	return &Storage{
		logger:          logger,
		agentController: agentController,
		dataStorage:     dataStorage,
		metricProvider:  metricProvider,

		tryReconnect: tryReconnect,

		storageMap:      make(map[uuid.UUID]rawFileStorageData, 10), //nolint:mnd // оптимизация
		storageMapMutex: &sync.RWMutex{},
	}
}

func (s *Storage) Init(ctx context.Context, skipNotAvailable bool) error {
	storages, err := s.dataStorage.FileStorages(ctx)
	if err != nil {
		return fmt.Errorf("get fs from db: %w", err)
	}

	for _, fs := range storages {
		if skipNotAvailable && fs.NotAvailable() {
			continue
		}

		storage, err := s.connect(ctx, fs)
		if err != nil {
			return fmt.Errorf("connect fs (%s): %w", fs.ID.String(), err)
		}

		// В данном случае делаем внесение данных по одному т.к.
		// снижение скорости инициализации из-за постоянных
		// блокировок некритично.
		s.storageMapMutex.Lock()
		s.storageMap[fs.ID] = storage
		s.storageMapMutex.Unlock()
	}

	return nil
}

func (s *Storage) FSIDForDownload(ctx context.Context) (uuid.UUID, error) {
	storages, err := s.dataStorage.FileStorages(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get fs from db: %w", err)
	}

	if len(storages) == 0 {
		return uuid.Nil, core.MissingFSError
	}

	slices.SortFunc(storages, func(a, b fsmodel.FileStorageSystem) int {
		return b.DownloadPriority - a.DownloadPriority
	})

	return storages[0].ID, nil
}
