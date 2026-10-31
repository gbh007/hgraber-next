package agent

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

func (uc *UseCase) CheckPages(ctx context.Context, pages []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error) {
	result := make([]entities.AgentPageCheckResult, len(pages))

	for i, p := range pages {
		hasParser, err := uc.loader.HasParser(ctx, p.BookURL)

		switch {
		case err != nil:
			result[i] = entities.AgentPageCheckResult{
				BookURL:     p.BookURL,
				ImageURL:    p.ImageURL,
				HasError:    true,
				ErrorReason: err.Error(),
			}

		case hasParser:
			result[i] = entities.AgentPageCheckResult{
				BookURL:    p.BookURL,
				ImageURL:   p.ImageURL,
				IsPossible: true,
			}

		default:
			result[i] = entities.AgentPageCheckResult{
				BookURL:       p.BookURL,
				ImageURL:      p.ImageURL,
				IsUnsupported: true,
			}
		}
	}

	return result, nil
}
