package agentparseruc

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

func (uc *UseCase) CheckBooks(ctx context.Context, urls []url.URL) ([]serveragentdto.AgentBookCheckResult, error) {
	result := make([]serveragentdto.AgentBookCheckResult, 0, len(urls))

	for _, u := range urls {
		hasParser, err := uc.loader.HasParser(ctx, u)
		if err != nil {
			result = append(result, serveragentdto.AgentBookCheckResult{
				URL:         u,
				HasError:    true,
				ErrorReason: err.Error(),
			})

			continue
		}

		if !hasParser {
			result = append(result, serveragentdto.AgentBookCheckResult{
				URL:           u,
				IsUnsupported: true,
			})

			continue
		}

		singleResult := serveragentdto.AgentBookCheckResult{
			URL:        u,
			IsPossible: true,
		}

		result = append(result, singleResult)
	}

	return result, nil
}
