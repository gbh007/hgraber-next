package parsingusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
)

func (uc *UseCase) PagesExists(
	ctx context.Context,
	urls []agentmodel.AgentPageURL,
) ([]agentmodel.AgentPageCheckResult, error) {
	result := make([]agentmodel.AgentPageCheckResult, 0, len(urls))

urlLoop:
	for _, u := range urls {
		pages, err := uc.storage.PagesByURL(ctx, u.ImageURL)
		if err != nil {
			return nil, fmt.Errorf("get pages by url (%s): %w", u.ImageURL.String(), err)
		}

		for _, p := range pages {
			if p.IsLoaded() {
				result = append(result, agentmodel.AgentPageCheckResult{
					BookURL:    u.BookURL,
					ImageURL:   u.ImageURL,
					IsPossible: true,
				})

				continue urlLoop
			}
		}

		result = append(result, agentmodel.AgentPageCheckResult{
			BookURL:       u.BookURL,
			ImageURL:      u.ImageURL,
			IsUnsupported: true,
		})
	}

	return result, nil
}
