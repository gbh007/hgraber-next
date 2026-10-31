package parsingusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

func (uc *UseCase) PagesExists(
	ctx context.Context,
	urls []serveragentdto.AgentPageURL,
) ([]serveragentdto.AgentPageCheckResult, error) {
	result := make([]serveragentdto.AgentPageCheckResult, 0, len(urls))

urlLoop:
	for _, u := range urls {
		pages, err := uc.storage.PagesByURL(ctx, u.ImageURL)
		if err != nil {
			return nil, fmt.Errorf("get pages by url (%s): %w", u.ImageURL.String(), err)
		}

		for _, p := range pages {
			if p.IsLoaded() {
				result = append(result, serveragentdto.AgentPageCheckResult{
					BookURL:    u.BookURL,
					ImageURL:   u.ImageURL,
					IsPossible: true,
				})

				continue urlLoop
			}
		}

		result = append(result, serveragentdto.AgentPageCheckResult{
			BookURL:       u.BookURL,
			ImageURL:      u.ImageURL,
			IsUnsupported: true,
		})
	}

	return result, nil
}
