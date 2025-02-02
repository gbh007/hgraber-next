package parsing

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) BooksExists(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error) {
	result := make([]entities.AgentBookCheckResult, 0, len(urls))

urlLoop:
	for _, u := range urls {
		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return nil, fmt.Errorf("get books by url (%s): %w", u.String(), err)
		}

		for _, id := range ids {
			book, err := uc.bookRequester.BookOriginFull(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("get book (%s) details by url (%s): %w", id.String(), u.String(), err)
			}

			// Только загруженные книги считаем доступными.
			if book.IsLoaded() {
				result = append(result, entities.AgentBookCheckResult{
					URL:        u,
					IsPossible: true,
				})

				continue urlLoop
			}
		}

		result = append(result, entities.AgentBookCheckResult{
			URL:           u,
			IsUnsupported: true,
		})
	}

	return result, nil
}
