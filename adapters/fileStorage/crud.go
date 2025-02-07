package fileStorage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (s *Storage) Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("create", &fsID, time.Since(startAt))
	}()

	if fsID == uuid.Nil {
		if s.legacyFileStorage == nil {
			return fmt.Errorf("%w: legacy", core.MissingFSError)
		}

		return s.legacyFileStorage.FS.Create(ctx, fileID, body)
	}

	storage, err := s.getFS(ctx, fsID, s.tryReconnect)
	if err != nil {
		return fmt.Errorf("get fs: %w", err)
	}

	return storage.FS.Create(ctx, fileID, body)
}

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) error {
	targetFSID, err := s.searchFS(ctx, fileID, fsID)
	if err != nil {
		return fmt.Errorf("search fs id: %w", err)
	}

	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("delete", &targetFSID, time.Since(startAt))
	}()

	if targetFSID == uuid.Nil {
		if s.legacyFileStorage == nil {
			return fmt.Errorf("%w: legacy", core.MissingFSError)
		}

		return s.legacyFileStorage.FS.Delete(ctx, fileID)
	}

	storage, err := s.getFS(ctx, targetFSID, s.tryReconnect)
	if err != nil {
		return fmt.Errorf("get fs: %w", err)
	}

	return storage.FS.Delete(ctx, fileID)
}

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error) {
	targetFSID, err := s.searchFS(ctx, fileID, fsID)
	if err != nil {
		return nil, fmt.Errorf("search fs id: %w", err)
	}

	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("get", &targetFSID, time.Since(startAt))
	}()

	if targetFSID == uuid.Nil {
		if s.legacyFileStorage == nil {
			return nil, fmt.Errorf("%w: legacy", core.MissingFSError)
		}

		return s.legacyFileStorage.FS.Get(ctx, fileID)
	}

	storage, err := s.getFS(ctx, targetFSID, s.tryReconnect)
	if err != nil {
		return nil, fmt.Errorf("get fs: %w", err)
	}

	return storage.FS.Get(ctx, fileID)
}

func (s *Storage) State(ctx context.Context, includeFileIDs bool, includeFileSizes bool, fsID uuid.UUID) (fsmodel.FSState, error) {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("state", &fsID, time.Since(startAt))
	}()

	if fsID == uuid.Nil {
		if s.legacyFileStorage == nil {
			return fsmodel.FSState{}, fmt.Errorf("%w: legacy", core.MissingFSError)
		}

		return s.legacyFileStorage.FS.State(ctx, includeFileIDs, includeFileSizes)
	}

	storage, err := s.getFS(ctx, fsID, s.tryReconnect)
	if err != nil {
		return fsmodel.FSState{}, fmt.Errorf("get fs: %w", err)
	}

	return storage.FS.State(ctx, includeFileIDs, includeFileSizes)
}
