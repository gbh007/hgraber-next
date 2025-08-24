package fileStorage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/agentFS"
	"github.com/gbh007/hgraber-next/adapters/localFiles"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (s *Storage) searchFS(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (uuid.UUID, error) {
	startAt := time.Now()

	defer func() {
		s.metricProvider.RegisterFSActionTime("search_fs", fsID, time.Since(startAt))
	}()

	if fsID != nil {
		return *fsID, nil
	}

	file, err := s.dataStorage.File(ctx, fileID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get file from db: %w", err)
	}

	return file.FSID, nil
}

func (s *Storage) getFS(ctx context.Context, fsID uuid.UUID, tryReconnect bool) (rawFileStorageData, error) {
	startAt := time.Now()

	defer func() {
		s.metricProvider.RegisterFSActionTime("get_fs", &fsID, time.Since(startAt))
	}()

	s.storageMapMutex.RLock()
	storage, ok := s.storageMap[fsID]
	s.storageMapMutex.RUnlock()

	switch {
	case ok:
		return storage, nil

	case !ok && tryReconnect:
		fsInfo, err := s.dataStorage.FileStorage(ctx, fsID)
		if err != nil {
			return rawFileStorageData{}, fmt.Errorf("get fs from db: %w", err)
		}

		storage, err = s.connect(ctx, fsInfo)
		if err != nil {
			return rawFileStorageData{}, fmt.Errorf("connect fs: %w", err)
		}

		s.storageMapMutex.Lock()
		s.storageMap[fsID] = storage
		s.storageMapMutex.Unlock()

		return storage, nil
	}

	return rawFileStorageData{}, core.MissingFSError
}

func (s *Storage) FSChange(ctx context.Context, fsID uuid.UUID, deleted bool) error {
	startAt := time.Now()

	defer func() {
		s.metricProvider.RegisterFSActionTime("change_fs", &fsID, time.Since(startAt))
	}()

	s.storageMapMutex.Lock()
	defer s.storageMapMutex.Unlock()

	delete(s.storageMap, fsID)

	if deleted {
		return nil
	}

	fsInfo, err := s.dataStorage.FileStorage(ctx, fsID)
	if err != nil {
		return fmt.Errorf("get fs from db: %w", err)
	}

	storage, err := s.connect(ctx, fsInfo)
	if err != nil {
		return fmt.Errorf("connect fs: %w", err)
	}

	s.storageMap[fsID] = storage

	return nil
}

func (s *Storage) connect(_ context.Context, fs fsmodel.FileStorageSystem) (rawFileStorageData, error) {
	var (
		err     error
		storage rawFileStorage
	)

	switch {
	case fs.AgentID != uuid.Nil:
		storage = agentFS.New(fs.AgentID, s.logger, s.agentController)

		raw := rawFileStorageData{
			FS:      storage,
			AgentID: fs.AgentID,
		}

		if fs.HighwayEnabled && fs.HighwayAddr != nil {
			raw.EnableHighway = true
			raw.HighwayServerScheme = fs.HighwayAddr.Scheme
			raw.HighwayServerHostWithPort = fs.HighwayAddr.Host
		}

		return raw, nil

	case fs.Path != "":
		storage, err = localFiles.New(fs.Path, s.logger)
		if err != nil {
			return rawFileStorageData{}, fmt.Errorf("fail init local file storage: %w", err)
		}

		return rawFileStorageData{
			FS:   storage,
			Path: fs.Path,
		}, nil

	default:
		return rawFileStorageData{}, fmt.Errorf("no configuration for file storage")
	}
}
