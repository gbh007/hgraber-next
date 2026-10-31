package agent

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

func (uc *UseCase) CheckBooks(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	result := make([]entities.AgentBookCheckResult, 0, len(urls))

	for _, u := range urls {
		hasParser, err := uc.loader.HasParser(ctx, u)
		if err != nil {
			result = append(result, entities.AgentBookCheckResult{
				URL:         u,
				HasError:    true,
				ErrorReason: err.Error(),
			})

			continue
		}

		if !hasParser {
			result = append(result, entities.AgentBookCheckResult{
				URL:           u,
				IsUnsupported: true,
			})

			continue
		}

		singleResult := entities.AgentBookCheckResult{
			URL:        u,
			IsPossible: true,
		}

		result = append(result, singleResult)
	}

	return result, nil
}
