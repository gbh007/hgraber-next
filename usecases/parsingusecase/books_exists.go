package parsingusecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
)

// TODO: тут вообще по факту не модели агента должны быть
func (uc *UseCase) BooksExists(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error) {
	result := make([]agentmodel.AgentBookCheckResult, 0, len(urls))

urlLoop:
	for _, u := range urls {
		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return nil, fmt.Errorf("get books by url (%s): %w", u.String(), err)
		}

		for _, id := range ids {
			book, err := uc.bookAdapter.BookRaw(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("get book (%s) details by url (%s): %w", id.String(), u.String(), err)
			}

			// Только загруженные книги считаем доступными.
			if book.IsLoaded() {
				result = append(result, agentmodel.AgentBookCheckResult{
					URL:        u,
					IsPossible: true,
				})

				continue urlLoop
			}
		}

		result = append(result, agentmodel.AgentBookCheckResult{
			URL:           u,
			IsUnsupported: true,
		})
	}

	return result, nil
}
