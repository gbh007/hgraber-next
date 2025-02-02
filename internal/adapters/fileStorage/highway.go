package fileStorage

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (s *Storage) HighwayFileURL(ctx context.Context, fileID uuid.UUID, ext string, fsID uuid.UUID) (url.URL, bool, error) {
	if fsID == uuid.Nil { // Легаси система не поддерживает highway
		return url.URL{}, false, nil
	}

	info, err := s.getFS(ctx, fsID, s.tryReconnect)
	if err != nil {
		return url.URL{}, false, fmt.Errorf("get fs for highway url: %w", err)
	}

	if !info.EnableHighway {
		return url.URL{}, false, nil
	}

	if info.HighwayTokenValidUntil.Before(time.Now()) {
		// TODO: возможно не самое подходящее место
		info.HighwayToken, err = s.refreshHighwayToken(ctx, fsID)
		if err != nil {
			return url.URL{}, false, fmt.Errorf("refresh token: %w", err)
		}
	}

	u := url.URL{
		Scheme: info.HighwayServerScheme,
		Host:   info.HighwayServerHostWithPort,
		Path:   "/api/highway/file/" + fileID.String() + url.PathEscape(ext),
	}

	v := url.Values{}
	v.Add("token", info.HighwayToken)
	u.RawQuery = v.Encode()

	return u, true, nil
}

func (s *Storage) refreshHighwayToken(ctx context.Context, fsID uuid.UUID) (string, error) {
	s.storageMapMutex.RLock()
	storage, ok := s.storageMap[fsID]
	s.storageMapMutex.RUnlock()

	if !ok {
		return "", entities.MissingFSError
	}

	if storage.AgentID == uuid.Nil {
		return "", fmt.Errorf("can't refresh token without agent fs")
	}

	token, validUntil, err := s.agentController.CreateHighwayToken(ctx, storage.AgentID)
	if err != nil {
		return "", fmt.Errorf("api: create highway token: %w", err)
	}

	s.storageMapMutex.Lock()
	storage = s.storageMap[fsID]
	storage.HighwayToken = token
	storage.HighwayTokenValidUntil = validUntil
	s.storageMap[fsID] = storage
	s.storageMapMutex.Unlock()

	return token, nil
}
