package fileStorage

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (s *Storage) FSList(ctx context.Context) ([]entities.FSWithStatus, error) {
	storages, err := s.dataStorage.FileStorages(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fs from db: %w", err)
	}

	result := make([]entities.FSWithStatus, 0, len(storages)+1)

	for _, storage := range storages {
		result = append(result, entities.FSWithStatus{
			Info: storage,
		})
	}

	if s.legacyFileStorage != nil {
		ls := entities.FSWithStatus{
			Info: entities.FileStorageSystem{
				Name:        "legacy storage",
				Description: "Устаревшее хранилище, крайне рекомендуется перейти на множественные файловые системы",
				AgentID:     s.legacyFileStorage.AgentID,
				Path:        s.legacyFileStorage.Path,
			},
			IsLegacy: true,
		}

		result = append(result, ls)
	}

	return result, nil
}
