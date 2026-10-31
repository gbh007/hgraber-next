package agentparseruc

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

func (uc *UseCase) CheckPages(ctx context.Context, pages []serveragentdto.AgentPageURL) ([]serveragentdto.AgentPageCheckResult, error) {
	result := make([]serveragentdto.AgentPageCheckResult, len(pages))

	for i, p := range pages {
		hasParser, err := uc.loader.HasParser(ctx, p.BookURL)

		switch {
		case err != nil:
			result[i] = serveragentdto.AgentPageCheckResult{
				BookURL:     p.BookURL,
				ImageURL:    p.ImageURL,
				HasError:    true,
				ErrorReason: err.Error(),
			}

		case hasParser:
			result[i] = serveragentdto.AgentPageCheckResult{
				BookURL:    p.BookURL,
				ImageURL:   p.ImageURL,
				IsPossible: true,
			}

		default:
			result[i] = serveragentdto.AgentPageCheckResult{
				BookURL:       p.BookURL,
				ImageURL:      p.ImageURL,
				IsUnsupported: true,
			}
		}
	}

	return result, nil
}
