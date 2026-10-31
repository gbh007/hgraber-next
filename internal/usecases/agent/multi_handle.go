package agent

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

func (uc *UseCase) MultiHandle(ctx context.Context, multiUrl url.URL) ([]serveragentdto.AgentBookCheckResult, error) {
	result := make([]serveragentdto.AgentBookCheckResult, 0, 100)

	urls, err := uc.loader.AllBooks(ctx, multiUrl)
	// FIXME: подумать как лучше это сделать
	if errors.Is(err, parsing.InvalidLinkError) {
		return []serveragentdto.AgentBookCheckResult{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("load all books: %w", err)
	}

	for _, stringURL := range urls {
		u, err := url.Parse(stringURL)
		if err != nil {
			return nil, fmt.Errorf("url parse (%s): %w", stringURL, err)
		}

		singleResult := serveragentdto.AgentBookCheckResult{
			URL:        *u,
			IsPossible: true,
		}

		result = append(result, singleResult)
	}

	return result, nil
}
