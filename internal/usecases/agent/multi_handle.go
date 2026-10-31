package agent

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/hgraber"
	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

func (uc *UseCase) MultiHandle(ctx context.Context, multiUrl url.URL) ([]entities.AgentBookCheckResult, error) {
	result := make([]entities.AgentBookCheckResult, 0, 100)

	urls, err := uc.loader.AllBooks(ctx, multiUrl)
	// FIXME: подумать как лучше это сделать
	if errors.Is(err, hgraber.InvalidLinkError) {
		return []entities.AgentBookCheckResult{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("load all books: %w", err)
	}

	for _, stringURL := range urls {
		u, err := url.Parse(stringURL)
		if err != nil {
			return nil, fmt.Errorf("url parse (%s): %w", stringURL, err)
		}

		singleResult := entities.AgentBookCheckResult{
			URL:        *u,
			IsPossible: true,
		}

		result = append(result, singleResult)
	}

	return result, nil
}
